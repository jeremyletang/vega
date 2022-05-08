package validators_test

import (
	"bytes"
	"context"
	"encoding/hex"
	"testing"

	"code.vegaprotocol.io/vega/validators"

	abcitypes "github.com/tendermint/tendermint/abci/types"
	types1 "github.com/tendermint/tendermint/proto/tendermint/types"

	commandspb "code.vegaprotocol.io/protos/vega/commands/v1"
	snapshot "code.vegaprotocol.io/protos/vega/snapshot/v1"
	vegactx "code.vegaprotocol.io/vega/libs/context"
	"code.vegaprotocol.io/vega/types"

	"code.vegaprotocol.io/vega/libs/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var topKey = (&types.PayloadTopology{}).Key()

func TestEmptySnapshot(t *testing.T) {
	top := getTestTopology(t)
	defer top.ctrl.Finish()

	s, p, err := top.GetState(topKey)
	assert.Nil(t, err)
	assert.Empty(t, p)
	assert.NotEmpty(t, s)

	assert.Equal(t, 1, len(top.Keys()))
}

func TestChangeOnValidatorPerfUpdate(t *testing.T) {
	top := getTestTopology(t)
	defer top.ctrl.Finish()

	s, _, err := top.GetState(topKey)
	assert.Nil(t, err)
	assert.NotEmpty(t, s)

	updateValidatorPerformanceToNonDefaultState(t, top.Topology)

	s2, _, err := top.GetState(topKey)
	assert.Nil(t, err)
	assert.NotEmpty(t, s2)
	require.False(t, bytes.Equal(s, s2))
}

func TestTopologySnapshot(t *testing.T) {
	top := getTestTopWithDefaultValidator(t)
	updateValidatorPerformanceToNonDefaultState(t, top.Topology)
	defer top.ctrl.Finish()

	s1, _, err := top.GetState(topKey)
	require.Nil(t, err)

	tmPubKeys := []string{"2w5hxsVqWFTV6/f0swyNVqOhY1vWI42MrfO0xkUqsiA=", "67g7+123M0kfMR35U7LLq09eEU1dVr6jHBEgEtPzkrs="}
	ctx := context.Background()

	nr1 := commandspb.AnnounceNode{
		Id:              "vega-master-pubkey",
		ChainPubKey:     tmPubKeys[0],
		VegaPubKey:      hex.EncodeToString([]byte("vega-key")),
		EthereumAddress: "eth-address",
	}
	err = top.AddNewNode(ctx, &nr1, validators.ValidatorStatusTendermint)
	assert.NoError(t, err)

	nr2 := commandspb.AnnounceNode{
		Id:              "vega-master-pubkey-2",
		ChainPubKey:     tmPubKeys[1],
		VegaPubKey:      hex.EncodeToString([]byte("vega-key-2")),
		EthereumAddress: "eth-address-2",
	}
	err = top.AddNewNode(ctx, &nr2, validators.ValidatorStatusTendermint)
	assert.NoError(t, err)

	kr1 := &commandspb.KeyRotateSubmission{
		NewPubKeyIndex:    1,
		TargetBlock:       10,
		NewPubKey:         "new-vega-key",
		CurrentPubKeyHash: hashKey("vega-key"),
	}
	err = top.AddKeyRotate(ctx, nr1.Id, 5, kr1)
	assert.NoError(t, err)

	kr2 := &commandspb.KeyRotateSubmission{
		NewPubKeyIndex:    1,
		TargetBlock:       11,
		NewPubKey:         "new-vega-key-2",
		CurrentPubKeyHash: hashKey("vega-key-2"),
	}
	err = top.AddKeyRotate(ctx, nr2.Id, 5, kr2)
	assert.NoError(t, err)

	ekr1 := &commandspb.EthereumKeyRotateSubmission{
		TargetBlock:    10,
		CurrentAddress: "eth-address",
		NewAddress:     "0x69bA3B3e6B5b1226A2e26De9a9E2D9C98f2b144B",
	}
	err = top.RotateEthereumKey(ctx, nr1.Id, 5, ekr1)
	assert.NoError(t, err)

	ekr2 := &commandspb.EthereumKeyRotateSubmission{
		TargetBlock:    11,
		CurrentAddress: "eth-address-2",
		NewAddress:     "0xd6B6e9514f2793Af89745Fd69FDa0DAbC228d336",
	}
	err = top.RotateEthereumKey(ctx, nr2.Id, 5, ekr2)
	assert.NoError(t, err)

	// Check the hashes have changed after each state change
	s3, _, err := top.GetState(topKey)
	require.Nil(t, err)
	require.False(t, bytes.Equal(s1, s3))

	// Get the state ready to load into a new instance of the engine
	state, _, _ := top.GetState(topKey)
	snap := &snapshot.Payload{}
	err = proto.Unmarshal(state, snap)
	require.Nil(t, err)

	snapTop := getTestTopWithDefaultValidator(t)
	defer snapTop.ctrl.Finish()

	ctx = vegactx.WithBlockHeight(ctx, 100)
	_, err = snapTop.LoadState(ctx, types.PayloadFromProto(snap))
	require.Nil(t, err)

	// Check the new reloaded engine is the same as the original
	s4, _, err := snapTop.GetState(topKey)
	require.Nil(t, err)
	require.True(t, bytes.Equal(s3, s4))
	assert.ElementsMatch(t, top.AllNodeIDs(), snapTop.AllNodeIDs())
	assert.ElementsMatch(t, top.AllVegaPubKeys(), snapTop.AllVegaPubKeys())
	assert.Equal(t, top.IsValidator(), snapTop.IsValidator())
	assert.Equal(t, top.GetPendingKeyRotation(kr1.TargetBlock, nr1.Id), snapTop.GetPendingKeyRotation(kr1.TargetBlock, nr1.Id))
	assert.Equal(t, top.GetPendingKeyRotation(kr2.TargetBlock, nr2.Id), snapTop.GetPendingKeyRotation(kr2.TargetBlock, nr2.Id))
	assert.Equal(t, top.GetPendingEthereumKeyRotation(ekr1.TargetBlock, nr1.Id), snapTop.GetPendingEthereumKeyRotation(ekr1.TargetBlock, nr1.Id))
	assert.Equal(t, top.GetPendingEthereumKeyRotation(ekr2.TargetBlock, nr2.Id), snapTop.GetPendingEthereumKeyRotation(ekr2.TargetBlock, nr2.Id))
}

func updateValidatorPerformanceToNonDefaultState(t *testing.T, top *validators.Topology) {
	t.Helper()
	req1 := abcitypes.RequestBeginBlock{Header: types1.Header{ProposerAddress: address1, Height: int64(1)}}
	top.BeginBlock(context.Background(), req1)

	// expecting address1 to propose but got address3
	req2 := abcitypes.RequestBeginBlock{Header: types1.Header{ProposerAddress: address3, Height: int64(1)}}
	top.BeginBlock(context.Background(), req2)
}
