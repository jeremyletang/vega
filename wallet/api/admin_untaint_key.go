package api

import (
	"context"
	"fmt"

	"code.vegaprotocol.io/vega/libs/jsonrpc"
	"github.com/mitchellh/mapstructure"
)

type AdminUntaintKeyParams struct {
	Wallet     string `json:"wallet"`
	PublicKey  string `json:"publicKey"`
	Passphrase string `json:"passphrase"`
}

type AdminUntaintKey struct {
	walletStore WalletStore
}

// Handle marks the specified public key as tainted. It makes it unusable for
// transaction signing.
func (h *AdminUntaintKey) Handle(ctx context.Context, rawParams jsonrpc.Params) (jsonrpc.Result, *jsonrpc.ErrorDetails) {
	params, err := validateUntaintKeyParams(rawParams)
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

	if !w.HasPublicKey(params.PublicKey) {
		return nil, invalidParams(ErrPublicKeyDoesNotExist)
	}

	if err := w.UntaintKey(params.PublicKey); err != nil {
		return nil, internalError(fmt.Errorf("could not remove the taint from the key: %w", err))
	}

	if err := h.walletStore.SaveWallet(ctx, w, params.Passphrase); err != nil {
		return nil, internalError(fmt.Errorf("could not save the wallet: %w", err))
	}

	return nil, nil
}

func validateUntaintKeyParams(rawParams jsonrpc.Params) (AdminUntaintKeyParams, error) {
	if rawParams == nil {
		return AdminUntaintKeyParams{}, ErrParamsRequired
	}

	params := AdminUntaintKeyParams{}
	if err := mapstructure.Decode(rawParams, &params); err != nil {
		return AdminUntaintKeyParams{}, ErrParamsDoNotMatch
	}

	if params.Wallet == "" {
		return AdminUntaintKeyParams{}, ErrWalletIsRequired
	}

	if params.Passphrase == "" {
		return AdminUntaintKeyParams{}, ErrPassphraseIsRequired
	}

	if params.PublicKey == "" {
		return AdminUntaintKeyParams{}, ErrPublicKeyIsRequired
	}

	return params, nil
}

func NewAdminUntaintKey(
	walletStore WalletStore,
) *AdminUntaintKey {
	return &AdminUntaintKey{
		walletStore: walletStore,
	}
}
