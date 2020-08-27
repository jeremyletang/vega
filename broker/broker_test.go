// +build !race

package broker_test

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"

	"code.vegaprotocol.io/vega/broker"
	"code.vegaprotocol.io/vega/broker/mocks"
	"code.vegaprotocol.io/vega/events"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type brokerTst struct {
	*broker.Broker
	cfunc context.CancelFunc
	ctx   context.Context
	ctrl  *gomock.Controller
}

type evt struct {
	t   events.Type
	ctx context.Context
	id  string
}

func getBroker(t *testing.T) *brokerTst {
	ctx, cfunc := context.WithCancel(context.Background())
	ctrl := gomock.NewController(t)
	return &brokerTst{
		Broker: broker.New(ctx),
		cfunc:  cfunc,
		ctx:    ctx,
		ctrl:   ctrl,
	}
}

func (b brokerTst) randomEvt() *evt {
	idString := "generic-id"
	if ctxV, ok := b.ctx.Value("traceID").(string); ok {
		idString = ctxV
	}
	return &evt{
		t:   events.All,
		ctx: b.ctx,
		id:  idString,
	}
}

func (b *brokerTst) Finish() {
	b.cfunc()
	b.ctrl.Finish()
}

func TestSubscribe(t *testing.T) {
	t.Run("Subscribe and unsubscribe required - success", testSubUnsubSuccess)
	t.Run("Subscribe reuses keys", testSubReuseKey)
	t.Run("Unsubscribe automatically if subscriber is closed", testAutoUnsubscribe)
}

func TestSendEvent(t *testing.T) {
	t.Run("Skip optional subscribers", testSkipOptional)
	t.Run("Skip optional subscribers in a batch send", testSendBatchChannel)
	t.Run("Send batch to ack subscriber", testSendBatch)
	t.Run("Stop sending if context is cancelled", testStopCtx)
	t.Run("Skip subscriber based on channel state", testSubscriberSkip)
	t.Run("Send only to typed subscriber", testEventTypeSubscription)
}

func testSubUnsubSuccess(t *testing.T) {
	broker := getBroker(t)
	defer broker.Finish()
	sub := mocks.NewMockSubscriber(broker.ctrl)
	reqSub := mocks.NewMockSubscriber(broker.ctrl)
	// subscribe + unsubscribe -> 2 calls
	sub.EXPECT().Types().Times(2).Return(nil)
	sub.EXPECT().Ack().Times(1).Return(false)
	reqSub.EXPECT().Types().Times(2).Return(nil)
	reqSub.EXPECT().Ack().Times(1).Return(true)
	k1 := broker.Subscribe(sub)    // not required
	k2 := broker.Subscribe(reqSub) // required
	assert.NotZero(t, k1)
	assert.NotZero(t, k2)
	assert.NotEqual(t, k1, k2)
	broker.Unsubscribe(k1)
	broker.Unsubscribe(k2)
	// no calls to subs expected once they are unsubscribed
	broker.Send(broker.randomEvt())
}

func testSubReuseKey(t *testing.T) {
	broker := getBroker(t)
	defer broker.Finish()
	sub := mocks.NewMockSubscriber(broker.ctrl)
	sub.EXPECT().Types().Times(4).Return(nil)
	sub.EXPECT().Ack().Times(1).Return(false)
	k1 := broker.Subscribe(sub)
	sub.EXPECT().Ack().Times(1).Return(true)
	assert.NotZero(t, k1)
	broker.Unsubscribe(k1)
	k2 := broker.Subscribe(sub)
	assert.Equal(t, k1, k2)
	broker.Unsubscribe(k2)
	// second unsubscribe is a no-op
	broker.Unsubscribe(k1)
}

func testAutoUnsubscribe(t *testing.T) {
	broker := getBroker(t)
	defer broker.Finish()
	sub := mocks.NewMockSubscriber(broker.ctrl)
	// sub, auto-unsub, sub again
	sub.EXPECT().Types().Times(3).Return(nil)
	sub.EXPECT().Ack().Times(1).Return(true)
	k1 := broker.Subscribe(sub)
	assert.NotZero(t, k1)
	// set up sub to be closed
	skipCh := make(chan struct{})
	closedCh := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(1)
	defer func() {
		close(skipCh)
	}()
	close(closedCh) // close the closed channel, so the subscriber is marked as closed when we try to send an event
	sub.EXPECT().Skip().AnyTimes().Return(skipCh)
	sub.EXPECT().Closed().AnyTimes().Return(closedCh).Do(func() {
		// indicator this function has been called already
		wg.Done()
	})
	// send an event, the subscriber should be marked as closed, and automatically unsubscribed
	broker.Send(broker.randomEvt())
	// introduce some wait mechanism here, because the unsubscribe call acquires its own lock now
	// so it's possible we haven't unsubscribed yet... the waitgroup should introduce enough time
	wg.Wait()
	// now try and subscribe again, the key should be reused
	sub.EXPECT().Ack().Times(1).Return(false)
	k2 := broker.Subscribe(sub)
	assert.Equal(t, k1, k2)
}

func testSendBatch(t *testing.T) {
	tstBroker := getBroker(t)
	sub := mocks.NewMockSubscriber(tstBroker.ctrl)
	cancelCh := make(chan struct{})
	defer func() {
		tstBroker.Finish()
		close(cancelCh)
	}()
	sub.EXPECT().Types().Times(1).Return(nil)
	sub.EXPECT().Ack().AnyTimes().Return(true)
	k1 := tstBroker.Subscribe(sub)
	assert.NotZero(t, k1)
	data := []events.Event{
		tstBroker.randomEvt(),
		tstBroker.randomEvt(),
		tstBroker.randomEvt(),
	}
	// ensure all 3 events are being sent (wait for routine to spawn)
	sub.EXPECT().Closed().AnyTimes().Return(cancelCh)
	sub.EXPECT().Skip().AnyTimes().Return(cancelCh)
	wg := sync.WaitGroup{}
	wg.Add(1)
	sub.EXPECT().Push(gomock.Any()).Times(1).Do(func(evts ...events.Event) {
		assert.Equal(t, len(data), len(evts))
		wg.Done()
	})

	// send events
	tstBroker.SendBatch(data)
	wg.Wait()
}

func testSendBatchChannel(t *testing.T) {
	tstBroker := getBroker(t)
	sub := mocks.NewMockSubscriber(tstBroker.ctrl)
	skipCh, closedCh, cCh := make(chan struct{}), make(chan struct{}), make(chan events.Event, 1)
	defer func() {
		tstBroker.Finish()
		close(closedCh)
		close(skipCh)
	}()
	twg := sync.WaitGroup{}
	twg.Add(2)
	sub.EXPECT().Types().Times(2).Return(nil).Do(func() {
		twg.Done()
	})
	sub.EXPECT().Ack().AnyTimes().Return(false)
	k1 := tstBroker.Subscribe(sub)
	assert.NotZero(t, k1)
	events := []events.Event{
		tstBroker.randomEvt(),
		tstBroker.randomEvt(),
		tstBroker.randomEvt(),
	}
	// ensure all 3 events are being sent (wait for routine to spawn)
	wg := sync.WaitGroup{}
	wg.Add(len(events))
	sub.EXPECT().Closed().AnyTimes().Return(closedCh)
	sub.EXPECT().Skip().AnyTimes().Return(skipCh)
	// we try to get the channel 3 times, only 1 of the attempts will actually publish the event
	sub.EXPECT().C().Times(len(events)).Return(cCh).Do(func() {
		// Done call each time we tried sending an event
		wg.Done()
	})

	// send events
	tstBroker.SendBatch(events)
	wg.Wait()
	// we've tried to send 3 events, subscriber could only accept one. Check state of all the things
	// we need to unsubscribe the subscriber, because we're closing the channels and race detector complains
	// because there's a loop calling functions that are returning the channels we're closing here
	tstBroker.Unsubscribe(k1)
	// ensure unsubscribe has returned
	twg.Wait()
	assert.Equal(t, events[0], <-cCh)
	// make sure the channel is empty (no writes were pending)
	assert.Equal(t, 0, len(cCh))
}

func testSkipOptional(t *testing.T) {
	tstBroker := getBroker(t)
	sub := mocks.NewMockSubscriber(tstBroker.ctrl)
	skipCh, closedCh, cCh := make(chan struct{}), make(chan struct{}), make(chan events.Event, 1)
	defer func() {
		tstBroker.Finish()
		close(closedCh)
		close(skipCh)
	}()
	twg := sync.WaitGroup{}
	twg.Add(2)
	sub.EXPECT().Types().Times(2).Return(nil).Do(func() {
		twg.Done()
	})
	sub.EXPECT().Ack().AnyTimes().Return(false)
	k1 := tstBroker.Subscribe(sub)
	assert.NotZero(t, k1)

	events := []*evt{
		tstBroker.randomEvt(),
		tstBroker.randomEvt(),
		tstBroker.randomEvt(),
	}
	// ensure all 3 events are being sent (wait for routine to spawn)
	wg := sync.WaitGroup{}
	wg.Add(len(events))
	sub.EXPECT().Closed().AnyTimes().Return(closedCh)
	sub.EXPECT().Skip().AnyTimes().Return(skipCh)
	// we try to get the channel 3 times, only 1 of the attempts will actually publish the event
	sub.EXPECT().C().Times(len(events)).Return(cCh).Do(func() {
		// Done call each time we tried sending an event
		wg.Done()
	})

	// send events
	for _, e := range events {
		tstBroker.Send(e)
	}
	wg.Wait()
	// we've tried to send 3 events, subscriber could only accept one. Check state of all the things
	// we need to unsubscribe the subscriber, because we're closing the channels and race detector complains
	// because there's a loop calling functions that are returning the channels we're closing here
	tstBroker.Unsubscribe(k1)
	// ensure unsubscribe has returned
	twg.Wait()
	assert.Equal(t, events[0], <-cCh)
	// make sure the channel is empty (no writes were pending)
	assert.Equal(t, 0, len(cCh))
}

func testStopCtx(t *testing.T) {
	broker := getBroker(t)
	defer broker.Finish()
	sub := mocks.NewMockSubscriber(broker.ctrl)
	ch := make(chan struct{})
	sub.EXPECT().Closed().AnyTimes().Return(ch)
	sub.EXPECT().Skip().AnyTimes().Return(ch)
	// no calls sub are expected, we cancelled the context
	broker.cfunc()
	sub.EXPECT().Types().Times(2).Return(nil)
	sub.EXPECT().Ack().AnyTimes().Return(true)
	k1 := broker.Subscribe(sub) // required sub
	assert.NotZero(t, k1)
	broker.Send(broker.randomEvt())
	// calling unsubscribe acquires lock, so we can ensure the Send call has returned
	broker.Unsubscribe(k1)
	close(ch)
}

func testSubscriberSkip(t *testing.T) {
	broker := getBroker(t)
	defer broker.Finish()
	sub := mocks.NewMockSubscriber(broker.ctrl)
	skipCh, closeCh := make(chan struct{}), make(chan struct{})
	skip := int64(0)
	events := []*evt{
		broker.randomEvt(),
		broker.randomEvt(),
	}
	wg := sync.WaitGroup{}
	wg.Add(len(events))
	sub.EXPECT().Closed().AnyTimes().Return(closeCh).Do(func() {
		wg.Done()
	})
	sub.EXPECT().Skip().AnyTimes().DoAndReturn(func() <-chan struct{} {
		// ensure at least all events + 1 skip are called
		if s := atomic.AddInt64(&skip, 1); s == 1 {
			// skip the first one
			ch := make(chan struct{})
			// return closed channel, so this subscriber is marked to skip events
			close(ch)
			return ch
		}
		return skipCh
	})
	// we expect this call once, and only for the SECOND call
	sub.EXPECT().Push(events[1]).Times(1)
	sub.EXPECT().Types().Times(2).Return(nil)
	sub.EXPECT().Ack().AnyTimes().Return(true)
	k1 := broker.Subscribe(sub) // required sub
	assert.NotZero(t, k1)
	for _, e := range events {
		broker.Send(e)
	}
	wg.Wait()
	// calling unsubscribe acquires lock, so we can ensure the Send call has returned
	broker.Unsubscribe(k1)
	close(skipCh)
	close(closeCh)
}

// test making sure that events are sent only to subs that are interested in it
func testEventTypeSubscription(t *testing.T) {
	broker := getBroker(t)
	defer broker.Finish()
	sub := mocks.NewMockSubscriber(broker.ctrl)
	allSub := mocks.NewMockSubscriber(broker.ctrl)
	diffSub := mocks.NewMockSubscriber(broker.ctrl)
	skipCh, closeCh := make(chan struct{}), make(chan struct{})
	event := broker.randomEvt()
	event.t = events.TimeUpdate
	wg := sync.WaitGroup{}
	wg.Add(2)
	// Closed check
	sub.EXPECT().Closed().AnyTimes().Return(closeCh)
	diffSub.EXPECT().Closed().AnyTimes().Return(closeCh) // can use the same channels, we're not closing them anyway
	allSub.EXPECT().Closed().AnyTimes().Return(closeCh)
	// skip check
	sub.EXPECT().Skip().AnyTimes().Return(skipCh)
	allSub.EXPECT().Skip().AnyTimes().Return(skipCh)
	diffSub.EXPECT().Skip().AnyTimes().Return(skipCh)
	// actually push the event - diffSub expects nothing
	sub.EXPECT().Push(gomock.Any()).Times(1).Do(func(_ interface{}) {
		wg.Done()
	})
	allSub.EXPECT().Push(gomock.Any()).Times(1).Do(func(_ interface{}) {
		wg.Done()
	})
	// the event types this subscriber is interested in
	sub.EXPECT().Types().Times(2).Return([]events.Type{events.TimeUpdate})
	allSub.EXPECT().Types().Times(2).Return(nil) // subscribed to ALL events
	// fake type:
	different := events.Type(int(events.All) + int(events.TimeUpdate) + 1) // this value cannot exist as an events.Type value
	diffSub.EXPECT().Types().Times(2).Return([]events.Type{different})
	// subscribe the subscriberjk
	sub.EXPECT().Ack().AnyTimes().Return(true)
	diffSub.EXPECT().Ack().AnyTimes().Return(true)
	allSub.EXPECT().Ack().AnyTimes().Return(true)
	k1 := broker.Subscribe(sub)     // required sub
	k2 := broker.Subscribe(diffSub) // required sub, but won't be used anyway
	k3 := broker.Subscribe(allSub)
	assert.NotZero(t, k1)
	assert.NotZero(t, k2)
	assert.NotZero(t, k3)
	assert.NotEqual(t, k1, k2)
	// send the correct event
	broker.Send(event)
	// ensure the event was delivered
	wg.Wait()
	// unsubscribe the subscriber, now we're done
	broker.Unsubscribe(k1)
	broker.Unsubscribe(k2)
	broker.Unsubscribe(k3)
	close(skipCh)
	close(closeCh)
}

func (e evt) Type() events.Type {
	return e.t
}

func (e evt) Context() context.Context {
	return e.ctx
}

func (e evt) TraceID() string {
	return e.id
}
