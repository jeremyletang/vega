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
	infoLong = cli.LongDesc(`
		Get wallet information such as wallet ID, version and type.
	`)

	infoExample = cli.Examples(`
		# Get the wallet information
		{{.Software}} info --wallet WALLET
	`)
)

type GetInfoWalletHandler func(params api.AdminDescribeWalletParams) (api.AdminDescribeWalletResult, error)

func NewCmdGetInfoWallet(w io.Writer, rf *RootFlags) *cobra.Command {
	h := func(params api.AdminDescribeWalletParams) (api.AdminDescribeWalletResult, error) {
		s, err := wallets.InitialiseStore(rf.Home)
		if err != nil {
			return api.AdminDescribeWalletResult{}, fmt.Errorf("couldn't initialise wallets store: %w", err)
		}

		describeWallet := api.NewAdminDescribeWallet(s)
		rawResult, errorDetails := describeWallet.Handle(context.Background(), params)
		if errorDetails != nil {
			return api.AdminDescribeWalletResult{}, errors.New(errorDetails.Data)
		}
		return rawResult.(api.AdminDescribeWalletResult), nil
	}
	return BuildCmdGetInfoWallet(w, h, rf)
}

func BuildCmdGetInfoWallet(w io.Writer, handler GetInfoWalletHandler, rf *RootFlags) *cobra.Command {
	f := &GetWalletInfoFlags{}

	cmd := &cobra.Command{
		Use:     "info",
		Short:   "Get wallet information",
		Long:    infoLong,
		Example: infoExample,
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
				PrintGetWalletInfoResponse(w, resp)
			case flags.JSONOutput:
				return printer.FprintJSON(w, resp)
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&f.Wallet,
		"wallet", "w",
		"",
		"Name of the wallet to use",
	)
	cmd.Flags().StringVarP(&f.PassphraseFile,
		"passphrase-file", "p",
		"",
		"Path to the file containing the wallet's passphrase",
	)

	autoCompleteWallet(cmd, rf.Home)

	return cmd
}

type GetWalletInfoFlags struct {
	Wallet         string
	PassphraseFile string
}

func (f *GetWalletInfoFlags) Validate() (api.AdminDescribeWalletParams, error) {
	req := api.AdminDescribeWalletParams{}

	if len(f.Wallet) == 0 {
		return api.AdminDescribeWalletParams{}, flags.MustBeSpecifiedError("wallet")
	}
	req.Wallet = f.Wallet

	passphrase, err := flags.GetPassphrase(f.PassphraseFile)
	if err != nil {
		return api.AdminDescribeWalletParams{}, err
	}
	req.Passphrase = passphrase

	return req, nil
}

func PrintGetWalletInfoResponse(w io.Writer, resp api.AdminDescribeWalletResult) {
	p := printer.NewInteractivePrinter(w)

	str := p.String()
	defer p.Print(str)

	str.Text("Type:").NextLine().WarningText(resp.Type).NextLine()
	str.Text("Version:").NextLine().WarningText(fmt.Sprintf("%d", resp.Version)).NextLine()
	str.Text("ID:").NextLine().WarningText(resp.ID).NextLine()
}
