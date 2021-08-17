package checkpoint_test

import (
	"errors"
	"fmt"
	"testing"

	"code.vegaprotocol.io/protos/vega"
	"code.vegaprotocol.io/vega/checkpoint"
	"code.vegaprotocol.io/vega/checkpoint/mocks"
	"code.vegaprotocol.io/vega/crypto"
	"code.vegaprotocol.io/vega/types"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

type testEngine struct {
	*checkpoint.Engine
	ctrl *gomock.Controller
}

func getTestEngine(t *testing.T) *testEngine {
	ctrl := gomock.NewController(t)
	eng, _ := checkpoint.New()
	return &testEngine{
		Engine: eng,
		ctrl:   ctrl,
	}
}

func TestGetCheckpoints(t *testing.T) {
	t.Run("test getting checkpoints loading in components via constructor - no duplicates", testGetCheckpointsConstructor)
	t.Run("test getting checkpoints loading in components using Add method - no duplicates", testGetCheckpointsAdd)
	t.Run("test adding duplicate components using Add methods", testAddDuplicate)
	t.Run("test adding duplicate component via constructor", testDuplicateConstructor)
	t.Run("test getting checkpoints - error", testGetCheckpointsErr)
}

func TestLoadCheckpoints(t *testing.T) {
	t.Run("test loading checkpoints after generating them - success", testLoadCheckpoints)
	t.Run("load non-registered components", testLoadMissingCheckpoint)
	t.Run("load checkpoint with invalid hash", testLoadInvalidHash)
	t.Run("load sparse checkpoint", testLoadSparse)
	t.Run("error loading checkpoint", testLoadError)
}

func testGetCheckpointsConstructor(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	components := map[types.CheckpointName]*mocks.MockState{
		types.GovernanceCheckpoint: mocks.NewMockState(ctrl),
		types.AssetsCheckpoint:     mocks.NewMockState(ctrl),
	}
	for k, c := range components {
		c.EXPECT().Name().Times(1).Return(k)
	}
	eng, err := checkpoint.New(components[types.GovernanceCheckpoint], components[types.AssetsCheckpoint])
	require.NoError(t, err)
	data := map[types.CheckpointName][]byte{
		types.GovernanceCheckpoint: []byte("foodata"),
		types.AssetsCheckpoint:     []byte("bardata"),
	}
	for k, c := range components {
		c.EXPECT().Checkpoint().Times(1).Return(data[k], nil)
	}
	raw, err := eng.Checkpoint()
	require.NoError(t, err)
	// now to check if the checkpoint contains the expected data
	for k, c := range components {
		c.EXPECT().Load(data[k]).Times(1).Return(nil)
	}
	require.NoError(t, eng.Load(raw))
}

func testGetCheckpointsAdd(t *testing.T) {
	t.Parallel()
	eng := getTestEngine(t)
	defer eng.ctrl.Finish()
	components := map[types.CheckpointName]*mocks.MockState{
		types.GovernanceCheckpoint: mocks.NewMockState(eng.ctrl),
		types.AssetsCheckpoint:     mocks.NewMockState(eng.ctrl),
	}
	data := map[types.CheckpointName][]byte{
		types.GovernanceCheckpoint: []byte("foodata"),
		types.AssetsCheckpoint:     []byte("bardata"),
	}
	for k, c := range components {
		c.EXPECT().Name().Times(1).Return(k)
	}
	require.NoError(t, eng.Add(components[types.GovernanceCheckpoint], components[types.AssetsCheckpoint]))
	for k, c := range components {
		c.EXPECT().Checkpoint().Times(1).Return(data[k], nil)
	}
	raw, err := eng.Checkpoint()
	require.NoError(t, err)
	// now to check if the checkpoint contains the expected data
	for k, c := range components {
		c.EXPECT().Load(data[k]).Times(1).Return(nil)
	}
	require.NoError(t, eng.Load(raw))
}

func testAddDuplicate(t *testing.T) {
	t.Parallel()
	eng := getTestEngine(t)
	defer eng.ctrl.Finish()
	comp := mocks.NewMockState(eng.ctrl)
	comp.EXPECT().Name().Times(2).Return(types.GovernanceCheckpoint)
	require.NoError(t, eng.Add(comp, comp)) // adding the exact same component (same ptr value)
	comp2 := mocks.NewMockState(eng.ctrl)
	comp2.EXPECT().Name().Times(1).Return(types.GovernanceCheckpoint)
	require.Error(t, eng.Add(comp2))
}

func testDuplicateConstructor(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	comp := mocks.NewMockState(ctrl)
	comp.EXPECT().Name().Times(3).Return(types.GovernanceCheckpoint)
	comp2 := mocks.NewMockState(ctrl)
	comp2.EXPECT().Name().Times(1).Return(types.GovernanceCheckpoint)
	// this is all good
	eng, err := checkpoint.New(comp, comp)
	require.NoError(t, err)
	require.NotNil(t, eng)
	eng, err = checkpoint.New(comp, comp2)
	require.Error(t, err)
	require.Nil(t, eng)
}

func testLoadCheckpoints(t *testing.T) {
	t.Parallel()
	eng := getTestEngine(t)
	defer eng.ctrl.Finish()
	components := map[types.CheckpointName]*mocks.MockState{
		types.GovernanceCheckpoint: mocks.NewMockState(eng.ctrl),
		types.AssetsCheckpoint:     mocks.NewMockState(eng.ctrl),
	}
	data := map[types.CheckpointName][]byte{
		types.GovernanceCheckpoint: []byte("foodata"),
		types.AssetsCheckpoint:     []byte("bardata"),
	}
	for k, c := range components {
		c.EXPECT().Name().Times(1).Return(k)
	}
	require.NoError(t, eng.Add(components[types.GovernanceCheckpoint], components[types.AssetsCheckpoint]))
	for k, c := range components {
		c.EXPECT().Checkpoint().Times(1).Return(data[k], nil)
	}
	snapshot, err := eng.Checkpoint()
	require.NoError(t, err)
	require.NotEmpty(t, snapshot)
	// create new components to load data in to
	wComps := map[types.CheckpointName]*wrappedMock{
		types.GovernanceCheckpoint: wrapMock(mocks.NewMockState(eng.ctrl)),
		types.AssetsCheckpoint:     wrapMock(mocks.NewMockState(eng.ctrl)),
	}
	for k, c := range wComps {
		c.EXPECT().Name().Times(1).Return(k)
		c.EXPECT().Load(data[k]).Times(1).Return(nil)
	}
	newEng, err := checkpoint.New(wComps[types.GovernanceCheckpoint], wComps[types.AssetsCheckpoint])
	require.NoError(t, err)
	require.NotNil(t, newEng)
	require.NoError(t, newEng.Load(snapshot))
	for k, exp := range data {
		wc := wComps[k]
		require.EqualValues(t, exp, wc.data)
	}
}

func testLoadMissingCheckpoint(t *testing.T) {
	t.Parallel()
	eng := getTestEngine(t)
	defer eng.ctrl.Finish()

	// create checkpoint data
	cp := types.Checkpoint{
		Assets: []byte("assets"),
	}
	b, err := vega.Marshal(cp.IntoProto())
	require.NoError(t, err)
	snap := &vega.Snapshot{
		State: b,
		Hash:  crypto.Hash(cp.HashBytes()),
	}
	data, err := vega.Marshal(snap)
	err = eng.Load(data)
	require.Error(t, err)
	require.Equal(t, checkpoint.ErrUnknownCheckpointName, err)
	// now try to tamper with the data itself in such a way that the has no longer matches:
	cp.Assets = []byte("foobar")
	b, err = vega.Marshal(cp.IntoProto())
	require.NoError(t, err)
	snap.State = b
	data, err = vega.Marshal(snap)
	err = eng.Load(data)
	require.Error(t, err)
	require.Equal(t, checkpoint.ErrSnapshotHashIncorrect, err)
}

func testLoadInvalidHash(t *testing.T) {
	t.Parallel()
	eng := getTestEngine(t)
	defer eng.ctrl.Finish()

	cp := types.Checkpoint{
		Assets: []byte("assets"),
	}
	snap := &vega.Snapshot{
		Hash: crypto.Hash(cp.HashBytes()),
	}
	// update data -> hash is invalid
	cp.Assets = []byte("foobar")
	b, err := vega.Marshal(cp.IntoProto())
	require.NoError(t, err)
	snap.State = b
	data, err := vega.Marshal(snap)
	err = eng.Load(data)
	require.Error(t, err)
	require.Equal(t, checkpoint.ErrSnapshotHashIncorrect, err)
}

func testLoadSparse(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	components := map[types.CheckpointName]*mocks.MockState{
		types.GovernanceCheckpoint: mocks.NewMockState(ctrl),
		types.AssetsCheckpoint:     mocks.NewMockState(ctrl),
	}
	for k, c := range components {
		c.EXPECT().Name().Times(1).Return(k)
	}
	eng, err := checkpoint.New(components[types.GovernanceCheckpoint])
	require.NoError(t, err)
	data := map[types.CheckpointName][]byte{
		types.GovernanceCheckpoint: []byte("foodata"),
	}
	c := components[types.GovernanceCheckpoint]
	c.EXPECT().Checkpoint().Times(1).Return(data[types.GovernanceCheckpoint], nil)
	snapshot, err := eng.Checkpoint()
	require.NoError(t, err)
	require.NoError(t, eng.Add(components[types.AssetsCheckpoint])) // load another component, not part of the checkpoints map
	c.EXPECT().Load(data[types.GovernanceCheckpoint]).Times(1).Return(nil)
	require.NoError(t, eng.Load(snapshot))
}

func testLoadError(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	components := map[types.CheckpointName]*mocks.MockState{
		types.GovernanceCheckpoint: mocks.NewMockState(ctrl),
		types.AssetsCheckpoint:     mocks.NewMockState(ctrl),
	}
	for k, c := range components {
		c.EXPECT().Name().Times(1).Return(k)
	}
	eng, err := checkpoint.New(components[types.GovernanceCheckpoint], components[types.AssetsCheckpoint])
	require.NoError(t, err)
	data := map[types.CheckpointName][]byte{
		types.GovernanceCheckpoint: []byte("foodata"),
		types.AssetsCheckpoint:     []byte("bardata"),
	}
	for k, c := range components {
		c.EXPECT().Checkpoint().Times(1).Return(data[k], nil)
	}
	ret := map[types.CheckpointName]error{
		types.GovernanceCheckpoint: errors.New("random error"),
		types.AssetsCheckpoint:     nil, // we always load checkpoints in order, so bar will go first, and should not return an error
	}
	checkpoints, err := eng.Checkpoint()
	require.NoError(t, err)
	for k, r := range ret {
		c := components[k]
		c.EXPECT().Load(data[k]).Times(1).Return(r)
	}
	err = eng.Load(checkpoints)
	require.Error(t, err)
	require.Equal(t, ret[types.GovernanceCheckpoint], err)
}

func testGetCheckpointsErr(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	components := map[types.CheckpointName]*mocks.MockState{
		types.GovernanceCheckpoint: mocks.NewMockState(ctrl),
		types.AssetsCheckpoint:     mocks.NewMockState(ctrl),
	}
	for k, c := range components {
		c.EXPECT().Name().Times(1).Return(k)
	}
	eng, err := checkpoint.New(components[types.GovernanceCheckpoint], components[types.AssetsCheckpoint])
	require.NoError(t, err)
	data := map[types.CheckpointName][]byte{
		types.GovernanceCheckpoint: nil,
		types.AssetsCheckpoint:     []byte("bardata"),
	}
	errs := map[types.CheckpointName]error{
		types.GovernanceCheckpoint: fmt.Errorf("random error"),
		types.AssetsCheckpoint:     nil,
	}
	for k, c := range components {
		c.EXPECT().Checkpoint().Times(1).Return(data[k], errs[k])
	}
	checkpoints, err := eng.Checkpoint()
	require.Nil(t, checkpoints)
	require.Error(t, err)
	require.Equal(t, errs[types.GovernanceCheckpoint], err)
}

type wrappedMock struct {
	*mocks.MockState
	data []byte
}

func wrapMock(m *mocks.MockState) *wrappedMock {
	return &wrappedMock{
		MockState: m,
	}
}

func (w *wrappedMock) Load(data []byte) error {
	w.data = data
	return w.MockState.Load(data)
}
