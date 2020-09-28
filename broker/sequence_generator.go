package broker

import (
	"sync"

	"code.vegaprotocol.io/vega/events"
)

type SeqEvent interface {
	events.Event
	SetSequenceID(s uint64)
	Sequence() uint64 // this isn't used apart from in tests
}

type gen struct {
	mu       sync.Mutex
	blockSeq map[string]uint64
	blocks   []string
}

func newGen() *gen {
	return &gen{
		blockSeq: map[string]uint64{},
		blocks:   []string{},
	}
}

// setSequence adds sequence ID to the event objects, returns the arguments because
// the events might be passed by value (interface values)
// returns the more restrictive event object - once seq ID is set, it should be treated as RO
func (g *gen) setSequence(evts ...events.Event) []events.Event {
	ln := uint64(len(evts))
	if ln == 0 {
		return nil
	}
	hash := evts[0].TraceID()
	g.mu.Lock()
	cur, ok := g.blockSeq[hash]
	if !ok {
		g.blocks = append(g.blocks, hash)
		cur = 1
		g.blockSeq[hash] = cur
		// if we're adding a new hash, check if we're up to 3, and remove it if needed
		defer g.cleanID()
	}
	// set sequence ID to the next sequence ID available
	g.blockSeq[hash] += ln
	g.mu.Unlock()
	ret := make([]events.Event, 0, len(evts))
	// create slice of ids
	for i := range evts {
		e, ok := evts[i].(SeqEvent)
		if !ok {
			continue
		}
		e.SetSequenceID(uint64(i) + cur)
		ret = append(ret, e)
	}
	return ret
}

func (g *gen) cleanID() {
	g.mu.Lock()
	if len(g.blocks) == 3 {
		delete(g.blockSeq, g.blocks[0])
		g.blocks = g.blocks[1:]
	}
	g.mu.Unlock()
}
