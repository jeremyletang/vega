package validators

import (
	"context"
	"errors"
	"sort"

	"github.com/golang/protobuf/proto"

	"code.vegaprotocol.io/vega/libs/crypto"
	"code.vegaprotocol.io/vega/types"
)

var (
	key = (&types.PayloadWitness{}).Key()

	hashKeys = []string{
		key,
	}

	ErrSnapshotKeyDoesNotExist = errors.New("unknown key for witness snapshot")
)

type witnessSnapshotState struct {
	changed    bool
	hash       []byte
	serialised []byte
}

func (w *Witness) Namespace() types.SnapshotNamespace {
	return types.WitnessSnapshot
}

func (w *Witness) Keys() []string {
	return hashKeys
}

func (w *Witness) serialise() ([]byte, error) {
	needResendRes := make([]string, 0, len(w.needResendRes))
	for r := range w.needResendRes {
		needResendRes = append(needResendRes, r)
	}
	sort.Strings(needResendRes)

	resources := make([]*types.Resource, 0, len(w.resources))
	for id, r := range w.resources {
		votes := make([]string, 0, len(r.votes))
		for v := range r.votes {
			votes = append(votes, v)
		}
		sort.Strings(votes)
		resources = append(resources, &types.Resource{
			ID:         id,
			CheckUntil: r.checkUntil,
			Votes:      votes,
			State:      r.state,
		})
	}

	payload := types.Payload{
		Data: &types.PayloadWitness{
			Witness: &types.Witness{
				Resources:           resources,
				NeedResendResources: needResendRes,
			},
		},
	}
	x := payload.IntoProto()
	return proto.Marshal(x)
}

// get the serialised form and hash of the given key.
func (w *Witness) getSerialisedAndHash(k string) ([]byte, []byte, error) {
	if k != key {
		return nil, nil, ErrSnapshotKeyDoesNotExist
	}

	if !w.wss.changed {
		return w.wss.serialised, w.wss.hash, nil
	}

	data, err := w.serialise()
	if err != nil {
		return nil, nil, err
	}

	hash := crypto.Hash(data)
	w.wss.serialised = data
	w.wss.hash = hash
	w.wss.changed = false
	return data, hash, nil
}

func (w *Witness) GetHash(k string) ([]byte, error) {
	_, hash, err := w.getSerialisedAndHash(k)
	return hash, err
}

func (w *Witness) GetState(k string) ([]byte, error) {
	state, _, err := w.getSerialisedAndHash(k)
	return state, err
}

func (w *Witness) LoadState(ctx context.Context, p *types.Payload) ([]types.StateProvider, error) {
	if w.Namespace() != p.Data.Namespace() {
		return nil, types.ErrInvalidSnapshotNamespace
	}
	// see what we're reloading
	switch pl := p.Data.(type) {
	case *types.PayloadWitness:
		return nil, w.restore(ctx, pl.Witness)
	default:
		return nil, types.ErrUnknownSnapshotType
	}
}

func (w *Witness) restore(ctx context.Context, witness *types.Witness) error {
	w.resources = map[string]*res{}
	w.needResendRes = map[string]struct{}{}

	for _, r := range witness.NeedResendResources {
		w.needResendRes[r] = struct{}{}
	}

	for _, r := range witness.Resources {
		w.resources[r.ID] = &res{
			checkUntil: r.CheckUntil,
			state:      r.State,
			votes:      map[string]struct{}{},
		}
		for _, v := range r.Votes {
			w.resources[r.ID].votes[v] = struct{}{}
		}
	}

	w.wss.changed = true
	return nil
}

func (w *Witness) RestoreResource(r Resource, cb func(interface{}, bool)) error {
	if _, ok := w.resources[r.GetID()]; !ok {
		return ErrInvalidResourceIDForNodeVote
	}

	res := w.resources[r.GetID()]
	res.cb = cb
	res.res = r
	ctx, cfunc := context.WithDeadline(context.Background(), res.checkUntil)
	res.cfunc = cfunc
	if w.top.IsValidator() {
		go w.start(ctx, res)
	}
	w.wss.changed = true
	return nil

}
