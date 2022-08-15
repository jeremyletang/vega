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

package visor

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strconv"
	"sync"
	"syscall"
	"time"

	"code.vegaprotocol.io/vega/core/types"
	"code.vegaprotocol.io/vega/logging"
	"code.vegaprotocol.io/vega/visor/config"
	"code.vegaprotocol.io/vega/visor/utils"

	"golang.org/x/sync/errgroup"
)

const snapshotBlockHeightFlagName = "--snapshot.load-from-block-height"

type BinariesRunner struct {
	mut         sync.RWMutex
	running     map[string]*exec.Cmd
	binsFolder  string
	log         *logging.Logger
	stopTimeout time.Duration
}

func NewBinariesRunner(log *logging.Logger, binsFolder string, stopTimeout time.Duration) *BinariesRunner {
	return &BinariesRunner{
		binsFolder:  binsFolder,
		running:     map[string]*exec.Cmd{},
		log:         log,
		stopTimeout: stopTimeout,
	}
}

func (r *BinariesRunner) runBinary(ctx context.Context, binary string, args []string) error {
	binPath := path.Join(r.binsFolder, binary)
	if err := utils.EnsureBinary(binPath); err != nil {
		return fmt.Errorf("failed to locate binary %s %v: %w", binPath, args, err)
	}

	cmd := exec.CommandContext(ctx, binPath, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	r.log.Debug("Starting binary",
		logging.String("binaryPath", binPath),
		logging.Strings("args", args),
	)

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to execute binary %s %v: %w", binPath, args, err)
	}

	// Ensures that if one binary failes all of them are killed
	go func() {
		<-ctx.Done()

		if cmd.Process == nil {
			return
		}

		// Process has already exited - no need to kill it
		if cmd.ProcessState != nil {
			return
		}

		r.log.Debug("Killing binary", logging.String("binaryPath", binPath))

		if err := cmd.Process.Kill(); err != nil {
			r.log.Debug("Failed to kill binary",
				logging.String("binaryPath", binPath),
				logging.Error(err),
			)
		}
	}()

	r.mut.Lock()
	r.running[binPath] = cmd
	r.mut.Unlock()

	defer func() {
		r.mut.Lock()
		delete(r.running, binPath)
		r.mut.Unlock()
	}()

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("failed to execute binary %s %v: %w", binPath, args, err)
	}

	return nil
}

func (r *BinariesRunner) Run(ctx context.Context, runConf *config.RunConfig, rInfo *types.ReleaseInfo) chan error {
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		// TODO consider moving this logic somewhere else
		args := Args(runConf.Vega.Binary.Args)
		if rInfo != nil {
			args.Set(snapshotBlockHeightFlagName, strconv.FormatUint(rInfo.UpgradeBlockHeight+1, 10))
		}

		return r.runBinary(ctx, runConf.Vega.Binary.Path, args)
	})

	if runConf.DataNode != nil {
		eg.Go(func() error {
			return r.runBinary(ctx, runConf.DataNode.Binary.Path, runConf.DataNode.Binary.Args)
		})
	}

	errChan := make(chan error)

	go func() {
		err := eg.Wait()
		if err != nil {
			errChan <- err
		}
	}()

	return errChan
}

func (r *BinariesRunner) signal(signal syscall.Signal) error {
	r.mut.RLock()
	defer r.mut.RUnlock()

	var err error
	for binName, c := range r.running {
		r.log.Info("Signaling process",
			logging.String("binaryName", binName),
			logging.String("signal", signal.String()),
		)

		err = c.Process.Signal(signal)
		if err != nil {
			r.log.Error("Failed to signal running binary",
				logging.String("binaryPath", c.Path),
				logging.Error(err),
			)
		}
	}

	return err
}

func (r *BinariesRunner) Stop() error {
	if err := r.signal(syscall.SIGTERM); err != nil {
		return err
	}

	r.mut.RLock()
	timeout := time.After(r.stopTimeout)
	r.mut.RUnlock()

	ticker := time.NewTicker(time.Second / 10)
	defer ticker.Stop()

	for {
		select {
		case <-timeout:
			return fmt.Errorf("failed to gratefully shut down processes: timed out")
		case <-ticker.C:
			r.mut.RLock()
			if len(r.running) == 0 {
				return nil
			}
			r.mut.RUnlock()
		}
	}
}

func (r *BinariesRunner) Kill() error {
	return r.signal(syscall.SIGKILL)
}
