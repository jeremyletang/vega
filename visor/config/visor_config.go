// Copyright (c) 2022 Gobalsky Labs Limited
//
// Use of this software is governed by the Business Source License included
// in the LICENSE file and at https://www.mariadb.com/bsl11.
//
// Change Date: 18 months from the later of the date of the first publicly
// available Distribution of this version of the repository, and 25 June 2022.
//
// On the date above, in accordance with the Business Source License, use
// of this software will be governed by version 3 or later of the GNU General
// Public License.

package config

import (
	"context"
	"fmt"
	"path"
	"runtime"
	"sync"
	"time"

	"code.vegaprotocol.io/vega/logging"
	"code.vegaprotocol.io/vega/paths"
	"github.com/fsnotify/fsnotify"
	"golang.org/x/sync/errgroup"
)

const (
	currentFolder      = "current"
	genesisFolder      = "genesis"
	configFileName     = "config.toml"
	RunConfigFileName  = "run-config.toml"
	VegaBinaryName     = "vega"
	DataNodeBinaryName = "data-node"
)

type AssetsConfig struct {
	Vega     string  `toml:"vega"`
	DataNode *string `toml:"data_node"`
}

func (ac AssetsConfig) ToSlice() []string {
	s := []string{ac.Vega}
	if ac.DataNode != nil {
		s = append(s, *ac.DataNode)
	}
	return s
}

type AutoInstallConfig struct {
	Enabled               bool         `toml:"enabled"`
	GithubRepositoryOwner string       `toml:"repositoryOwner"`
	GithubRepository      string       `toml:"repository"`
	Assets                AssetsConfig `toml:"assets"`
}

type VisorConfigFile struct {
	UpgradeFolders           map[string]string `toml:"upgradeFolders"`
	MaxNumberOfRestarts      int               `toml:"maxNumberOfRestarts"`
	RestartsDelaySeconds     int               `toml:"restartsDelaySeconds"`
	StopSignalTimeoutSeconds int               `toml:"stopSignalTimeoutSeconds"`

	AutoInstall AutoInstallConfig `toml:"autoInstall"`
}

func parseAndValidateVisorConfigFile(path string) (*VisorConfigFile, error) {
	conf := VisorConfigFile{}
	if err := paths.ReadStructuredFile(path, &conf); err != nil {
		return nil, fmt.Errorf("failed to parse VisorConfig: %w", err)
	}

	return &conf, nil
}

type VisorConfig struct {
	mut        sync.RWMutex
	configPath string
	homePath   string
	data       *VisorConfigFile
	log        *logging.Logger
}

func DefaultVisorConfig(log *logging.Logger, homePath string) *VisorConfig {
	return &VisorConfig{
		log:        log,
		homePath:   homePath,
		configPath: path.Join(homePath, configFileName),
		data: &VisorConfigFile{
			UpgradeFolders:           map[string]string{"vX.X.X": "vX.X.X"},
			MaxNumberOfRestarts:      3,
			RestartsDelaySeconds:     5,
			StopSignalTimeoutSeconds: 15,
			AutoInstall: AutoInstallConfig{
				Enabled:               true,
				GithubRepositoryOwner: "vegaprotocol",
				GithubRepository:      "vega",
				Assets: AssetsConfig{
					Vega: fmt.Sprintf("vega-%s-%s", runtime.GOOS, "amd64"),
				},
			},
		},
	}
}

func NewVisorConfig(log *logging.Logger, homePath string) (*VisorConfig, error) {
	configPath := path.Join(homePath, configFileName)

	dataFile, err := parseAndValidateVisorConfigFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &VisorConfig{
		configPath: configPath,
		homePath:   homePath,
		data:       dataFile,
		log:        log,
	}, nil
}

func (pc *VisorConfig) reload() error {
	pc.log.Info("Reloading config")
	dataFile, err := parseAndValidateVisorConfigFile(pc.configPath)
	if err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}

	pc.mut.Lock()
	pc.data.UpgradeFolders = dataFile.UpgradeFolders
	pc.data.MaxNumberOfRestarts = dataFile.MaxNumberOfRestarts
	pc.data.RestartsDelaySeconds = dataFile.RestartsDelaySeconds
	pc.mut.Unlock()

	pc.log.Info("Reloading config success")

	return nil
}

func (pc *VisorConfig) WatchForUpdate(ctx context.Context) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	var eg errgroup.Group
	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case event, ok := <-watcher.Events:
				if !ok {
					return nil
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					// add a small sleep here in order to handle vi
					// vi do not send a write event / edit the file in place,
					// it always create a temporary file, then delete the original one,
					// and then rename the temp file with the name of the original file.
					// if we try to update the conf as soon as we get the event, the file is not
					// always created and we get a no such file or directory error
					time.Sleep(50 * time.Millisecond)

					pc.reload()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return nil
				}
				return err
			}
		}
	})

	if err := watcher.Add(pc.configPath); err != nil {
		return err
	}

	return eg.Wait()
}

func (pc *VisorConfig) CurrentFolder() string {
	pc.mut.RLock()
	defer pc.mut.RUnlock()

	return path.Join(pc.homePath, currentFolder)
}

func (pc *VisorConfig) CurrentRunConfigPath() string {
	pc.mut.RLock()
	defer pc.mut.RUnlock()

	return path.Join(pc.CurrentFolder(), RunConfigFileName)
}

func (pc *VisorConfig) GenesisFolder() string {
	pc.mut.RLock()
	defer pc.mut.RUnlock()

	return path.Join(pc.homePath, genesisFolder)
}

func (pc *VisorConfig) UpgradeFolder(releaseTag string) string {
	pc.mut.RLock()
	defer pc.mut.RUnlock()

	if folderName, ok := pc.data.UpgradeFolders[releaseTag]; ok {
		return path.Join(pc.homePath, folderName)
	}

	return path.Join(pc.homePath, releaseTag)
}

func (pc *VisorConfig) MaxNumberOfRestarts() int {
	pc.mut.RLock()
	defer pc.mut.RUnlock()

	return pc.data.MaxNumberOfRestarts
}

func (pc *VisorConfig) RestartsDelaySeconds() int {
	pc.mut.RLock()
	defer pc.mut.RUnlock()

	return pc.data.RestartsDelaySeconds
}

func (pc *VisorConfig) StopSignalTimeoutSeconds() int {
	pc.mut.RLock()
	defer pc.mut.RUnlock()

	return pc.data.StopSignalTimeoutSeconds
}

func (pc *VisorConfig) AutoInstall() AutoInstallConfig {
	pc.mut.RLock()
	defer pc.mut.RUnlock()

	return pc.data.AutoInstall
}

func (pc *VisorConfig) WriteToFile() error {
	pc.mut.RLock()
	defer pc.mut.RUnlock()

	return paths.WriteStructuredFile(pc.configPath, pc.data)
}
