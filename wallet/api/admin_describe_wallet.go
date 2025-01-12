package api

import (
	"context"
	"fmt"

	"code.vegaprotocol.io/vega/libs/jsonrpc"
	"github.com/mitchellh/mapstructure"
)

type AdminDescribeWalletParams struct {
	Wallet     string `json:"wallet"`
	Passphrase string `json:"passphrase"`
}

type AdminDescribeWalletResult struct {
	Name    string `json:"name"`
	ID      string `json:"id"`
	Type    string `json:"type"`
	Version uint32 `json:"version"`
}

type AdminDescribeWallet struct {
	walletStore WalletStore
}

// Handle retrieve a wallet from its name and passphrase.
func (h *AdminDescribeWallet) Handle(ctx context.Context, rawParams jsonrpc.Params) (jsonrpc.Result, *jsonrpc.ErrorDetails) {
	params, err := validateDescribeWalletParams(rawParams)
	if err != nil {
		return nil, invalidParams(err)
	}

	if exist, err := h.walletStore.WalletExists(ctx, params.Wallet); err != nil {
		return nil, internalError(fmt.Errorf("could not verify the wallet existence: %w", err))
	} else if !exist {
		return nil, invalidParams(ErrWalletDoesNotExist)
	}

	w, err := h.walletStore.GetWallet(ctx, params.Wallet, params.Passphrase)
	if err != nil {
		return nil, internalError(fmt.Errorf("could not retrieve the wallet: %w", err))
	}

	return AdminDescribeWalletResult{
		Name:    w.Name(),
		ID:      w.ID(),
		Type:    w.Type(),
		Version: w.Version(),
	}, nil
}

func validateDescribeWalletParams(rawParams jsonrpc.Params) (AdminDescribeWalletParams, error) {
	if rawParams == nil {
		return AdminDescribeWalletParams{}, ErrParamsRequired
	}

	params := AdminDescribeWalletParams{}
	if err := mapstructure.Decode(rawParams, &params); err != nil {
		return AdminDescribeWalletParams{}, ErrParamsDoNotMatch
	}

	if params.Wallet == "" {
		return AdminDescribeWalletParams{}, ErrWalletIsRequired
	}

	if params.Passphrase == "" {
		return AdminDescribeWalletParams{}, ErrPassphraseIsRequired
	}

	return params, nil
}

func NewAdminDescribeWallet(
	walletStore WalletStore,
) *AdminDescribeWallet {
	return &AdminDescribeWallet{
		walletStore: walletStore,
	}
}
