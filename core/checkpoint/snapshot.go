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

package checkpoint

import (
	"context"
	"time"

	"code.vegaprotocol.io/vega/core/types"
	"code.vegaprotocol.io/vega/logging"

	"code.vegaprotocol.io/vega/libs/proto"
)

func (e *Engine) Namespace() types.SnapshotNamespace {
	return e.state.Namespace()
}

func (e *Engine) Keys() []string {
	return []string{
		e.state.Key(),
	}
}

func (e *Engine) Stopped() bool {
	return false
}

func (e *Engine) setNextCP(t time.Time) {
	e.nextCP = t
	e.state.Checkpoint.NextCp = t.UnixNano()
	e.data = []byte{}
	e.updated = true
}

func (e *Engine) GetState(k string) ([]byte, []types.StateProvider, error) {
	if k != e.state.Key() {
		return nil, nil, types.ErrSnapshotKeyDoesNotExist
	}
	if len(e.data) == 0 {
		if err := e.serialiseState(); err != nil {
			return nil, nil, err
		}
	}
	return e.data, nil, nil
}

func (e *Engine) serialiseState() error {
	e.log.Debug("serialising checkpoint", logging.Int64("nextcp", e.state.Checkpoint.NextCp))
	pl := types.Payload{
		Data: e.state,
	}
	data, err := proto.Marshal(pl.IntoProto())
	if err != nil {
		return err
	}

	e.data = data
	return nil
}

func (e *Engine) LoadState(_ context.Context, snap *types.Payload) ([]types.StateProvider, error) {
	if snap.Namespace() != e.state.Namespace() {
		return nil, types.ErrInvalidSnapshotNamespace
	}
	if snap.Key() != e.state.Key() {
		return nil, types.ErrSnapshotKeyDoesNotExist
	}
	state := snap.Data.(*types.PayloadCheckpoint)
	e.state = state
	e.setNextCP(time.Unix(0, state.Checkpoint.NextCp))
	e.log.Debug("loaded checkpoint snapshot", logging.Int64("nextcp", e.state.Checkpoint.NextCp))
	return nil, nil
}

func (e *Engine) PollChanges(ctx context.Context, k string, ch chan<- *types.Payload) {
	e.poll = make(chan struct{})
	defer func() {
		close(e.poll)
	}()
	if k != e.state.Key() {
		e.snapErr = types.ErrSnapshotKeyDoesNotExist
		ch <- nil
		return
	}
	if !e.updated {
		// nil on channel indicates no changes
		ch <- nil
		return
	}
	// create the payload object for snapshot
	pl := types.Payload{
		Data: &types.PayloadCheckpoint{
			Checkpoint: &types.CPState{
				NextCp: e.nextCP.UnixNano(),
			},
		},
	}
	select {
	case <-ctx.Done():
		e.snapErr = ctx.Err()
		return
	default:
		// send new update, flag as done
		ch <- &pl
		e.updated = false
	}
}

func (e *Engine) Sync() error {
	<-e.poll
	return e.Err()
}

func (e *Engine) Err() error {
	err := e.snapErr
	// remove error
	e.snapErr = nil
	return err
}
