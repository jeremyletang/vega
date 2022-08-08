// Copyright (c) 2022 Gobalsky Labs Limited
//
// Use of this software is governed by the Business Source License included
// in the LICENSE.VEGA file and at https://www.mariadb.com/bsl11.
//
// Change Date: 18 months from the later of the date of the first publicly
// available Distribution of this version of the repository, and 25 June 2022.
//
// On the date above, in accordance with the Business Source License, use
// of this software will be governed by version 3 or later of the GNU General
// Public License.

package banking

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"time"

	"code.vegaprotocol.io/vega/core/events"
	"code.vegaprotocol.io/vega/core/types"
	vgcrypto "code.vegaprotocol.io/vega/libs/crypto"
	"code.vegaprotocol.io/vega/libs/num"
	"code.vegaprotocol.io/vega/logging"
)

func (e *Engine) WithdrawBuiltinAsset(
	ctx context.Context, id, party, assetID string, amount *num.Uint,
) error {
	// build the withdrawal type
	w, ref := e.newWithdrawal(id, party, assetID, amount, time.Time{}, nil)
	w.Status = types.WithdrawalStatusRejected // default
	e.withdrawals[w.ID] = withdrawalRef{w, ref}

	asset, err := e.assets.Get(assetID)
	if err != nil {
		e.broker.Send(events.NewWithdrawalEvent(ctx, *w))
		e.log.Error("unable to get asset by id",
			logging.AssetID(assetID),
			logging.Error(err))
		return err
	}

	if !asset.IsBuiltinAsset() {
		e.broker.Send(events.NewWithdrawalEvent(ctx, *w))
		return ErrWrongAssetTypeUsedInBuiltinAssetChainEvent
	}

	return e.finalizeWithdraw(ctx, w)
}

func (e *Engine) DepositBuiltinAsset(
	ctx context.Context, d *types.BuiltinAssetDeposit, id string, nonce uint64,
) error {
	dep := e.newDeposit(id, d.PartyID, d.VegaAssetID, d.Amount, "") // no hash
	e.broker.Send(events.NewDepositEvent(ctx, *dep))
	asset, err := e.assets.Get(d.VegaAssetID)
	if err != nil {
		dep.Status = types.DepositStatusCancelled
		e.broker.Send(events.NewDepositEvent(ctx, *dep))
		e.log.Error("unable to get asset by id",
			logging.AssetID(d.VegaAssetID),
			logging.Error(err))
		return err
	}
	if !asset.IsBuiltinAsset() {
		dep.Status = types.DepositStatusCancelled
		e.broker.Send(events.NewDepositEvent(ctx, *dep))
		return ErrWrongAssetTypeUsedInBuiltinAssetChainEvent
	}

	// create a pretend "hash" from the nonce (which is randomly generated by the faucet)
	// ready for calls to getRef()
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, nonce)

	aa := &assetAction{
		id:       dep.ID,
		state:    pendingState,
		builtinD: d,
		asset:    asset,
		txHash:   hex.EncodeToString(vgcrypto.Hash(b)),
	}
	e.assetActs[aa.id] = aa
	e.deposits[dep.ID] = dep
	return e.witness.StartCheck(aa, e.onCheckDone, e.timeService.GetTimeNow().Add(defaultValidationDuration))
}

func (e *Engine) EnableBuiltinAsset(ctx context.Context, assetID string) error {
	return e.finalizeAssetList(ctx, assetID)
}
