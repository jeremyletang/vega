package cmd

import (
	"context"
	"errors"
	"fmt"
	"io"

	"code.vegaprotocol.io/vega/cmd/vegawallet/commands/cli"
	"code.vegaprotocol.io/vega/cmd/vegawallet/commands/flags"
	"code.vegaprotocol.io/vega/cmd/vegawallet/commands/printer"
	"code.vegaprotocol.io/vega/wallet/api"
	"code.vegaprotocol.io/vega/wallet/wallets"

	"github.com/spf13/cobra"
)

var (
	isolateKeyLong = cli.LongDesc(`
		Extract the specified key pair into an isolated wallet.

		An isolated wallet is a wallet that contains a single key pair and that
		has been stripped from its cryptographic node.

		Removing the cryptographic node from the wallet minimizes the impact of a
		stolen wallet, as it makes it impossible to retrieve or generate keys out
		of it.

		This creates a wallet that is only able to sign and verify transactions.

		This adds an extra layer of security.
	`)

	isolateKeyExample = cli.Examples(`
		# Isolate a key pair
		{{.Software}} key isolate --wallet WALLET --pubkey PUBKEY
	`)
)

type IsolateKeyHandler func(api.AdminIsolateKeyParams) (api.AdminIsolateKeyResult, error)

func NewCmdIsolateKey(w io.Writer, rf *RootFlags) *cobra.Command {
	h := func(params api.AdminIsolateKeyParams) (api.AdminIsolateKeyResult, error) {
		s, err := wallets.InitialiseStore(rf.Home)
		if err != nil {
			return api.AdminIsolateKeyResult{}, fmt.Errorf("could not initialise wallets store: %w", err)
		}

		isolateKey := api.NewAdminIsolateKey(s)
		rawResult, errDetails := isolateKey.Handle(context.Background(), params)
		if errDetails != nil {
			return api.AdminIsolateKeyResult{}, errors.New(errDetails.Data)
		}
		return rawResult.(api.AdminIsolateKeyResult), nil
	}

	return BuildCmdIsolateKey(w, h, rf)
}

func BuildCmdIsolateKey(w io.Writer, handler IsolateKeyHandler, rf *RootFlags) *cobra.Command {
	f := &IsolateKeyFlags{}

	cmd := &cobra.Command{
		Use:     "isolate",
		Short:   "Extract the specified key pair into an isolated wallet",
		Long:    isolateKeyLong,
		Example: isolateKeyExample,
		RunE: func(_ *cobra.Command, _ []string) error {
			req, err := f.Validate()
			if err != nil {
				return err
			}

			resp, err := handler(req)
			if err != nil {
				return err
			}

			switch rf.Output {
			case flags.InteractiveOutput:
				PrintIsolateKeyResponse(w, resp)
			case flags.JSONOutput:
				return printer.FprintJSON(w, resp)
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&f.Wallet,
		"wallet", "w",
		"",
		"Wallet holding the public key",
	)
	cmd.Flags().StringVarP(&f.PubKey,
		"pubkey", "k",
		"",
		"Public key to isolate (hex-encoded)",
	)
	cmd.Flags().StringVarP(&f.PassphraseFile,
		"passphrase-file", "p",
		"",
		"Path to the file containing the wallet's passphrase",
	)

	autoCompleteWallet(cmd, rf.Home)

	return cmd
}

type IsolateKeyFlags struct {
	Wallet         string
	PubKey         string
	PassphraseFile string
}

func (f *IsolateKeyFlags) Validate() (api.AdminIsolateKeyParams, error) {
	if len(f.Wallet) == 0 {
		return api.AdminIsolateKeyParams{}, flags.MustBeSpecifiedError("wallet")
	}

	if len(f.PubKey) == 0 {
		return api.AdminIsolateKeyParams{}, flags.MustBeSpecifiedError("pubkey")
	}

	passphrase, err := flags.GetPassphrase(f.PassphraseFile)
	if err != nil {
		return api.AdminIsolateKeyParams{}, err
	}

	return api.AdminIsolateKeyParams{
		Wallet:     f.Wallet,
		PublicKey:  f.PubKey,
		Passphrase: passphrase,
	}, nil
}

func PrintIsolateKeyResponse(w io.Writer, resp api.AdminIsolateKeyResult) {
	p := printer.NewInteractivePrinter(w)

	str := p.String()
	defer p.Print(str)
	str.CheckMark().Text("Key pair has been isolated in wallet ").Bold(resp.Wallet).Text(" at: ").SuccessText(resp.FilePath).NextLine()
	str.CheckMark().SuccessText("Key isolation succeeded").NextLine()
}
