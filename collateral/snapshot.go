package collateral

import (
	"context"
	"sort"

	"code.vegaprotocol.io/vega/events"
	"code.vegaprotocol.io/vega/libs/crypto"
	"code.vegaprotocol.io/vega/logging"
	"code.vegaprotocol.io/vega/types"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

type accState struct {
	accPL      types.PayloadCollateralAccounts
	assPL      types.PayloadCollateralAssets
	assets     map[string]types.Asset
	assetIDs   []string
	hashes     map[string][]byte
	updates    map[string]bool
	serialised map[string][]byte
	hashKeys   []string
}

var (
	ErrSnapshotKeyDoesNotExist  = errors.New("unknown key for collateral snapshot")
	ErrInvalidSnapshotNamespace = errors.New("invalid snapshot namespace")
	ErrUnknownSnapshotType      = errors.New("snapshot data type not known")
)

func (e *Engine) Namespace() types.SnapshotNamespace {
	return types.CollateralSnapshot
}

func (e *Engine) Keys() []string {
	return e.state.hashKeys
}

func (e *Engine) GetHash(k string) ([]byte, error) {
	return e.state.getHash(k)
}

func (e *Engine) GetState(k string) ([]byte, []types.StateProvider, error) {
	state, err := e.state.getState(k)
	return state, nil, err
}

func (e *Engine) LoadState(ctx context.Context, p *types.Payload) ([]types.StateProvider, error) {
	if e.Namespace() != p.Data.Namespace() {
		return nil, ErrInvalidSnapshotNamespace
	}
	// see what we're reloading
	switch pl := p.Data.(type) {
	case *types.PayloadCollateralAssets:
		err := e.restoreAssets(ctx, pl.CollateralAssets)
		return nil, err
	case *types.PayloadCollateralAccounts:
		err := e.restoreAccounts(ctx, pl.CollateralAccounts)
		return nil, err
	default:
		return nil, ErrUnknownSnapshotType
	}
}

func (e *Engine) restoreAccounts(ctx context.Context, accs *types.CollateralAccounts) error {
	e.log.Debug("restoring accounts snapshot", logging.Int("n_accounts", len(accs.Accounts)))

	evts := []events.Event{}
	pevts := []events.Event{}
	e.accs = make(map[string]*types.Account, len(accs.Accounts))
	e.partiesAccs = map[string]map[string]*types.Account{}
	e.hashableAccs = make([]*types.Account, 0, len(accs.Accounts))
	for _, acc := range accs.Accounts {
		e.accs[acc.ID] = acc
		if _, ok := e.partiesAccs[acc.Owner]; !ok {
			e.partiesAccs[acc.Owner] = map[string]*types.Account{}
		}
		e.partiesAccs[acc.Owner][acc.ID] = acc
		if acc.Type != types.AccountTypeExternal {
			e.hashableAccs = append(e.hashableAccs, acc)
			e.addAccountToHashableSlice(acc)
		}
		evts = append(evts, events.NewAccountEvent(ctx, *acc))

		if acc.Owner != systemOwner {
			pevts = append(pevts, events.NewPartyEvent(ctx, types.Party{Id: acc.Owner}))
		}
	}
	e.state.updateAccs(e.hashableAccs)
	e.broker.SendBatch(evts)
	e.broker.SendBatch(pevts)
	return nil
}

func (e *Engine) restoreAssets(ctx context.Context, assets *types.CollateralAssets) error {
	// @TODO the ID and name might not be the same, perhaps we need
	// to wrap the asset details to preserve that data
	e.log.Debug("restoring assets snapshot", logging.Int("n_assets", len(assets.Assets)))
	e.enabledAssets = make(map[string]types.Asset, len(assets.Assets))
	e.state.assetIDs = make([]string, 0, len(assets.Assets))
	e.state.assets = make(map[string]types.Asset, len(assets.Assets))
	evts := []events.Event{}
	for _, a := range assets.Assets {
		ast := types.Asset{
			ID:      a.ID,
			Details: a.Details,
		}
		e.enabledAssets[a.ID] = ast
		e.state.enableAsset(ast)
		evts = append(evts, events.NewAssetEvent(ctx, *a))
	}
	e.broker.SendBatch(evts)
	return nil
}

func newAccState() *accState {
	state := &accState{
		accPL: types.PayloadCollateralAccounts{
			CollateralAccounts: &types.CollateralAccounts{},
		},
		assPL: types.PayloadCollateralAssets{
			CollateralAssets: &types.CollateralAssets{},
		},
		assets:     map[string]types.Asset{},
		assetIDs:   []string{},
		hashes:     map[string][]byte{},
		updates:    map[string]bool{},
		serialised: map[string][]byte{},
	}
	state.hashKeys = []string{
		state.assPL.Key(),
		state.accPL.Key(),
	}
	for _, k := range state.hashKeys {
		state.hashes[k] = nil
		state.updates[k] = true
		state.serialised[k] = nil
	}
	return state
}

func (a *accState) enableAsset(asset types.Asset) {
	a.assets[asset.ID] = asset
	a.assetIDs = append(a.assetIDs, asset.ID)
	sort.Strings(a.assetIDs)
	a.updates[a.assPL.Key()] = true
}

func (a *accState) updateAccs(accs []*types.Account) {
	a.updates[a.accPL.Key()] = true
	a.accPL.CollateralAccounts.Accounts = accs[:]
}

func (a *accState) hashAssets() error {
	k := a.assPL.Key()
	if !a.updates[k] {
		return nil
	}
	assets := make([]*types.Asset, 0, len(a.assetIDs))
	for _, id := range a.assetIDs {
		ast := a.assets[id]
		assets = append(assets, &ast)
	}
	a.assPL.CollateralAssets.Assets = assets
	pl := types.Payload{
		Data: &a.assPL,
	}
	data, err := proto.Marshal(pl.IntoProto())
	if err != nil {
		return err
	}
	a.updates[k] = false
	a.hashes[k] = crypto.Hash(data)
	a.serialised[k] = data
	return nil
}

func (a *accState) hashAccounts() error {
	k := a.accPL.Key()
	if !a.updates[k] {
		return nil
	}
	// the account slice is already set, sorted and all
	pl := types.Payload{
		Data: &a.accPL,
	}
	data, err := proto.Marshal(pl.IntoProto())
	if err != nil {
		return err
	}
	a.serialised[k] = data
	a.hashes[k] = crypto.Hash(data)
	a.updates[k] = false
	return nil
}

func (a *accState) getState(k string) ([]byte, error) {
	update, exist := a.updates[k]
	if !exist {
		return nil, ErrSnapshotKeyDoesNotExist
	}
	if !update {
		h := a.serialised[k]
		return h, nil
	}
	if k == a.assPL.Key() {
		if err := a.hashAssets(); err != nil {
			return nil, err
		}
	} else if err := a.hashAccounts(); err != nil {
		return nil, err
	}
	h := a.serialised[k]
	return h, nil
}

func (a *accState) getHash(k string) ([]byte, error) {
	update, exist := a.updates[k]
	if !exist {
		return nil, ErrSnapshotKeyDoesNotExist
	}
	// we have a pending update
	if update {
		// hash whichever one we need to update
		if k == a.assPL.Key() {
			if err := a.hashAssets(); err != nil {
				return nil, err
			}
		} else if err := a.hashAccounts(); err != nil {
			return nil, err
		}
	}
	// fetch the new hash and return
	h := a.hashes[k]
	return h, nil
}
