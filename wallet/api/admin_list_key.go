package api

import (
	"context"
	"fmt"

	"code.vegaprotocol.io/vega/libs/jsonrpc"
	"github.com/mitchellh/mapstructure"
)

type AdminListKeysParams struct {
	Wallet     string `json:"wallet"`
	Passphrase string `json:"passphrase"`
}

type AdminListKeysResult struct {
	PublicKeys []AdminNamedPublicKey `json:"keys"`
}

type AdminNamedPublicKey struct {
	Name      string `json:"name"`
	PublicKey string `json:"publicKey"`
}

type AdminListKeys struct {
	walletStore WalletStore
}

// Handle list all the generated key from the specified wallet.
func (h *AdminListKeys) Handle(ctx context.Context, rawParams jsonrpc.Params) (jsonrpc.Result, *jsonrpc.ErrorDetails) {
	params, err := validateAdminListKeysParams(rawParams)
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

	publicKeys := w.ListPublicKeys()
	if err != nil {
		return nil, internalError(fmt.Errorf("could not list the keys: %w", err))
	}

	strPublicKeys := make([]AdminNamedPublicKey, 0, len(publicKeys))
	for _, publicKey := range publicKeys {
		strPublicKeys = append(strPublicKeys, AdminNamedPublicKey{
			Name:      publicKey.Name(),
			PublicKey: publicKey.Key(),
		})
	}

	return AdminListKeysResult{
		PublicKeys: strPublicKeys,
	}, nil
}

func validateAdminListKeysParams(rawParams jsonrpc.Params) (AdminListKeysParams, error) {
	if rawParams == nil {
		return AdminListKeysParams{}, ErrParamsRequired
	}

	params := AdminListKeysParams{}
	if err := mapstructure.Decode(rawParams, &params); err != nil {
		return AdminListKeysParams{}, ErrParamsDoNotMatch
	}

	if params.Wallet == "" {
		return AdminListKeysParams{}, ErrWalletIsRequired
	}

	if params.Passphrase == "" {
		return AdminListKeysParams{}, ErrPassphraseIsRequired
	}

	return params, nil
}

func NewAdminListKeys(
	walletStore WalletStore,
) *AdminListKeys {
	return &AdminListKeys{
		walletStore: walletStore,
	}
}
