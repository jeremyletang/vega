package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"code.vegaprotocol.io/vega/fsutil"
	"code.vegaprotocol.io/vega/logging"
	"code.vegaprotocol.io/vega/wallet"
	"code.vegaprotocol.io/vega/wallet/crypto"

	"github.com/spf13/cobra"
)

type walletCommand struct {
	command

	rootPath    string
	walletOwner string
	passphrase  string
	Log         *logging.Logger
}

func (ic *walletCommand) Init(c *Cli) {
	ic.cli = c
	ic.cmd = &cobra.Command{
		Use:   "wallet",
		Short: "The wallet subcommand",
		Long:  "Create and manage wallets",
	}

	genkey := &cobra.Command{
		Use:   "genkey",
		Short: "Generate a new keypair for a wallet",
		Long:  "Generate a new keypair for a wallet, this will implicitly generate a new wallet if none exist for the given name",
		RunE:  ic.GenKey,
	}
	genkey.Flags().StringVarP(&ic.rootPath, "root-path", "r", fsutil.DefaultVegaDir(), "Path of the root directory in which the configuration will be located")
	genkey.Flags().StringVarP(&ic.walletOwner, "name", "n", "", "Name of the wallet to use")
	genkey.Flags().StringVarP(&ic.passphrase, "passphrase", "p", "", "Passphrase to access the wallet")
	ic.cmd.AddCommand(genkey)

	list := &cobra.Command{
		Use:   "list",
		Short: "List keypairs of a wallet",
		Long:  "List all the keypairs for a given wallet",
		RunE:  ic.List,
	}
	list.Flags().StringVarP(&ic.rootPath, "root-path", "r", fsutil.DefaultVegaDir(), "Path of the root directory in which the configuration will be located")
	list.Flags().StringVarP(&ic.walletOwner, "name", "n", "", "Name of the wallet to use")
	list.Flags().StringVarP(&ic.passphrase, "passphrase", "p", "", "Passphrase to access the wallet")
	ic.cmd.AddCommand(list)
}

func (ic *walletCommand) GenKey(cmd *cobra.Command, args []string) error {
	if len(ic.walletOwner) <= 0 {
		return errors.New("wallet name is required")
	}
	if len(ic.passphrase) <= 0 {
		return errors.New("passphrase is required")
	}

	if ok, err := fsutil.PathExists(ic.rootPath); !ok {
		return fmt.Errorf("invalid root directory path: %v", err)
	}

	if err := wallet.EnsureBaseFolder(ic.rootPath); err != nil {
		return fmt.Errorf("unable to initialization root folder: %v", err)
	}

	_, err := wallet.Read(ic.rootPath, ic.walletOwner, ic.passphrase)
	if err != nil {
		if err != wallet.ErrWalletDoesNotExist {
			// this an invalid key, returning error
			return fmt.Errorf("unable to decrypt wallet: %v\n", err)
		}
		// wallet do not exit, let's try to create it
		_, err = wallet.Create(ic.rootPath, ic.walletOwner, ic.passphrase)
		if err != nil {
			return fmt.Errorf("unable to create wallet: %v", err)
		}
	}

	// at this point we have a valid wallet
	// let's generate the keypair
	// defaulting to ed25519 for now
	algo := crypto.NewEd25519()
	kp, err := wallet.GenKeypair(algo.Name())
	if err != nil {
		return fmt.Errorf("unable to generate new key pair: %v", err)
	}

	// now updating the wallet and saving it
	_, err = wallet.AddKeypair(kp, ic.rootPath, ic.walletOwner, ic.passphrase)
	if err != nil {
		return fmt.Errorf("unable to add keypair to wallet: %v", err)
	}

	// print the new keys for user info
	fmt.Printf("new generated keys:\n")
	fmt.Printf("public: 0x%v\n", kp.Pub)
	fmt.Printf("private: 0x%v\n", kp.Priv)

	return nil
}

func (ic *walletCommand) List(cmd *cobra.Command, args []string) error {
	if len(ic.walletOwner) <= 0 {
		return errors.New("wallet name is required")
	}
	if len(ic.passphrase) <= 0 {
		return errors.New("passphrase is required")
	}

	if ok, err := fsutil.PathExists(ic.rootPath); !ok {
		return fmt.Errorf("invalid root directory path: %v", err)
	}

	w, err := wallet.Read(ic.rootPath, ic.walletOwner, ic.passphrase)
	if err != nil {
		return fmt.Errorf("unable to decrypt wallet: %v\n", err)
	}

	buf, err := json.MarshalIndent(w, " ", " ")
	if err != nil {
		return fmt.Errorf("unable to indent message: %v", err)
	}

	// print the new keys for user info
	fmt.Printf("List of all your keypairs:\n")
	fmt.Printf("%v\n", string(buf))

	return nil
}
