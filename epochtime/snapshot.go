package epochtime

import (
	"errors"

	"code.vegaprotocol.io/vega/libs/crypto"

	"code.vegaprotocol.io/vega/types"
	"github.com/golang/protobuf/proto"
)

var (
	ErrSnapshotKeyDoesNotExist = errors.New("unknown key for epochtime snapshot")
)

type snapState struct {
	hash []byte
	data []byte
	pl   *types.EpochState
	t    *types.PayloadEpoch
}

func newSnapState() *snapState {
	state := &snapState{
		pl: &types.EpochState{},
		t:  &types.PayloadEpoch{},
	}

	state.t.EpochState = state.pl
	return state
}

func (ss *snapState) serialise() error {

	// Already done it
	if len(ss.data) != 0 {
		return nil
	}

	data, err := proto.Marshal(ss.pl.IntoProto())
	if err != nil {
		return err
	}

	ss.data = data
	ss.hash = crypto.Hash(data)
	return nil

}

func (s *Svc) setSnapshot() {

	// Get a fresh one
	s.state = newSnapState()
	s.state.pl = &types.EpochState{
		Seq:                  s.epoch.Seq,
		StartTime:            s.epoch.StartTime,
		ExpireTime:           s.epoch.ExpireTime,
		EndTime:              s.epoch.EndTime,
		Action:               s.epoch.Action,
		ReadyToStartNewEpoch: s.readyToEndEpoch,
		ReadyToEndEpoch:      s.readyToEndEpoch,
	}

}

func (s *Svc) Namespace() types.SnapshotNamespace {
	return s.state.t.Namespace()
}

func (s *Svc) Keys() []string {
	return []string{s.state.t.Key()}
}

func (s *Svc) GetHash(_ string) ([]byte, error) {
	if err := s.state.serialise(); err != nil {
		return nil, err
	}
	return s.state.hash, nil
}

func (s *Svc) Snapshot() (map[string][]byte, error) {
	if err := s.state.serialise(); err != nil {
		return nil, err
	}
	return map[string][]byte{s.state.t.Key(): s.state.data}, nil
}

func (s *Svc) GetState(_ string) ([]byte, error) {
	if err := s.state.serialise(); err != nil {
		return nil, err
	}
	return s.state.data, nil
}

func (s *Svc) LoadSnapshot(payload *types.PayloadEpoch) error {

	snap := payload.EpochState

	s.epoch = types.Epoch{
		Seq:        snap.Seq,
		StartTime:  snap.StartTime,
		ExpireTime: snap.ExpireTime,
		EndTime:    snap.EndTime,
		Action:     snap.Action,
	}

	s.readyToStartNewEpoch = snap.ReadyToStartNewEpoch
	s.readyToEndEpoch = snap.ReadyToEndEpoch
	s.length = s.epoch.ExpireTime.Sub(s.epoch.StartTime)
	s.setSnapshot()
	return nil
}
