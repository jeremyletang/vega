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

package notary_test

import (
	"context"
	"testing"
	"time"

	bmocks "code.vegaprotocol.io/vega/core/broker/mocks"
	"code.vegaprotocol.io/vega/core/notary"
	"code.vegaprotocol.io/vega/core/notary/mocks"
	"code.vegaprotocol.io/vega/core/types"
	"code.vegaprotocol.io/vega/logging"
	commandspb "code.vegaprotocol.io/vega/protos/vega/commands/v1"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type testNotary struct {
	*notary.SnapshotNotary
	// 	*notary.Notary
	ctrl   *gomock.Controller
	top    *mocks.MockValidatorTopology
	cmd    *mocks.MockCommander
	onTick func(context.Context, time.Time)
}

func getTestNotary(t *testing.T) *testNotary {
	t.Helper()
	ctrl := gomock.NewController(t)
	top := mocks.NewMockValidatorTopology(ctrl)
	broker := bmocks.NewMockBroker(ctrl)
	cmd := mocks.NewMockCommander(ctrl)
	broker.EXPECT().Send(gomock.Any()).AnyTimes()
	broker.EXPECT().SendBatch(gomock.Any()).AnyTimes()
	notr := notary.NewWithSnapshot(logging.NewTestLogger(), notary.NewDefaultConfig(), top, broker, cmd)
	return &testNotary{
		SnapshotNotary: notr,
		top:            top,
		ctrl:           ctrl,
		cmd:            cmd,
		onTick:         notr.OnTick,
	}
}

func TestNotary(t *testing.T) {
	t.Run("test add key for unknow resource - fail", testAddKeyForKOResource)
	t.Run("test add bad signature for known resource - success", testAddBadSignatureForOKResource)
	t.Run("test add key finalize all sig", testAddKeyFinalize)
	t.Run("test add key finalize all fails if sigs aren't tendermint validators", testAddKeyFinalizeFails)
}

func testAddKeyForKOResource(t *testing.T) {
	notr := getTestNotary(t)
	kind := types.NodeSignatureKindAssetNew
	resID := "resid"
	key := "123456"
	sig := []byte("123456")

	ns := commandspb.NodeSignature{
		Sig:  sig,
		Id:   resID,
		Kind: kind,
	}

	// first try to add a key for invalid resource
	err := notr.RegisterSignature(context.Background(), key, ns)
	assert.EqualError(t, err, notary.ErrUnknownResourceID.Error())

	// then try to start twice an aggregate
	notr.top.EXPECT().IsValidator().Times(1).Return(true)

	notr.StartAggregate(resID, kind, sig)
	assert.Panics(t, func() { notr.StartAggregate(resID, kind, sig) }, "expect to panic")
}

func testAddBadSignatureForOKResource(t *testing.T) {
	notr := getTestNotary(t)

	kind := types.NodeSignatureKindAssetNew
	resID := "resid"
	key := "123456"
	sig := []byte("123456")

	// start to aggregate, being a validator or not here doesn't matter
	notr.top.EXPECT().IsValidator().Times(1).Return(false)
	notr.StartAggregate(resID, kind, nil) // we send nil here if we are no validator

	ns := commandspb.NodeSignature{
		Sig:  sig,
		Id:   resID,
		Kind: kind,
	}

	// The signature we have received is not from a validator
	notr.top.EXPECT().IsValidatorVegaPubKey(gomock.Any()).AnyTimes().Return(false)

	err := notr.RegisterSignature(context.Background(), key, ns)
	assert.EqualError(t, err, notary.ErrNotAValidatorSignature.Error())
}

func testAddKeyFinalize(t *testing.T) {
	notr := getTestNotary(t)

	kind := types.NodeSignatureKindAssetNew
	resID := "resid"
	key := "123456"
	sig := []byte("123456")

	// add a valid node
	notr.top.EXPECT().Len().AnyTimes().Return(1)
	notr.top.EXPECT().IsValidatorVegaPubKey(gomock.Any()).AnyTimes().Return(true)
	notr.top.EXPECT().IsTendermintValidator(gomock.Any()).AnyTimes().Return(true)

	notr.top.EXPECT().IsValidator().Times(1).Return(true)
	notr.StartAggregate(resID, kind, sig)

	// expect command to be send on next on time update
	notr.cmd.EXPECT().Command(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(1)
	notr.onTick(context.Background(), time.Now())

	ns := commandspb.NodeSignature{
		Sig:  sig,
		Id:   resID,
		Kind: kind,
	}

	// first try to add a key for invalid resource
	notr.top.EXPECT().SelfVegaPubKey().Times(1).Return(key)
	err := notr.RegisterSignature(context.Background(), key, ns)
	assert.NoError(t, err, notary.ErrUnknownResourceID.Error())

	signatures, ok := notr.IsSigned(context.Background(), resID, kind)
	assert.True(t, ok)
	assert.Len(t, signatures, 1)
}

func testAddKeyFinalizeFails(t *testing.T) {
	notr := getTestNotary(t)

	kind := types.NodeSignatureKindAssetNew
	resID := "resid"
	key := "123456"
	sig := []byte("123456")

	// add a valid node
	notr.top.EXPECT().Len().AnyTimes().Return(1)
	notr.top.EXPECT().IsValidatorVegaPubKey(gomock.Any()).AnyTimes().Return(true)
	notr.top.EXPECT().IsTendermintValidator(gomock.Any()).AnyTimes().Return(false)

	notr.top.EXPECT().IsValidator().Times(1).Return(true)
	notr.StartAggregate(resID, kind, sig)

	// expect command to be send on next on time update
	notr.cmd.EXPECT().Command(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(1)
	notr.onTick(context.Background(), time.Now())

	ns := commandspb.NodeSignature{
		Sig:  sig,
		Id:   resID,
		Kind: kind,
	}

	// first try to add a key for invalid resource
	notr.top.EXPECT().SelfVegaPubKey().Times(1).Return(key)
	err := notr.RegisterSignature(context.Background(), key, ns)
	assert.NoError(t, err, notary.ErrUnknownResourceID.Error())

	signatures, ok := notr.IsSigned(context.Background(), resID, kind)
	assert.False(t, ok)
	assert.Len(t, signatures, 0) // no signatures because everyone that signed wasn't a Tendermint validator
}
