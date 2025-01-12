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

package stubs

import (
	"context"
	"math/rand"
	"sort"
	"strconv"
	"time"

	"code.vegaprotocol.io/vega/core/types/statevar"
)

var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

type StateVarStub struct {
	seq                 int
	currentTime         time.Time
	svs                 map[string]*sv
	eventTypeToStateVar map[statevar.EventType][]*sv
	rng                 *rand.Rand
	readyForTimeTrigger map[string]struct{}
	stateVarToNextCalc  map[string]time.Time
	updateFrequency     time.Duration
}

func NewStateVar() *StateVarStub {
	return &StateVarStub{
		seq:                 0,
		svs:                 map[string]*sv{},
		eventTypeToStateVar: map[statevar.EventType][]*sv{},
		rng:                 rand.New(rand.NewSource(0)),
		readyForTimeTrigger: map[string]struct{}{},
		stateVarToNextCalc:  map[string]time.Time{},
	}
}

type sv struct {
	ID               string
	asset            string
	market           string
	converter        statevar.Converter
	startCalculation func(string, statevar.FinaliseCalculation)
	trigger          []statevar.EventType
	result           func(context.Context, statevar.StateVariableResult) error
	eventID          string
}

func (e *StateVarStub) OnFloatingPointUpdatesDurationUpdate(ctx context.Context, updateFrequency time.Duration) error {
	e.updateFrequency = updateFrequency
	return nil
}

func (e *StateVarStub) UnregisterStateVariable(asset, market string) {
}

func (e *StateVarStub) RegisterStateVariable(asset, market, name string, converter statevar.Converter, startCalculation func(string, statevar.FinaliseCalculation), trigger []statevar.EventType, result func(context.Context, statevar.StateVariableResult) error) error {
	ID := asset + "_" + market + "_" + name + "_" + strconv.Itoa(e.seq)
	e.seq++
	e.svs[ID] = &sv{
		ID:               ID,
		asset:            asset,
		market:           market,
		converter:        converter,
		startCalculation: startCalculation,
		trigger:          trigger,
		result:           result,
	}
	for _, t := range trigger {
		if _, ok := e.eventTypeToStateVar[t]; !ok {
			e.eventTypeToStateVar[t] = []*sv{}
		}
		e.eventTypeToStateVar[t] = append(e.eventTypeToStateVar[t], e.svs[ID])
	}
	e.ReadyForTimeTrigger(asset, market)
	return nil
}

func (e *StateVarStub) NewEvent(asset, market string, eventType statevar.EventType) {
	eventID := e.generateID(asset, market)
	for _, s := range e.eventTypeToStateVar[eventType] {
		if s.market != market || s.asset != asset {
			continue
		}
		s.eventID = eventID
		s.startCalculation(eventID, s)
		if _, ok := e.stateVarToNextCalc[s.ID]; ok {
			e.stateVarToNextCalc[s.ID] = e.currentTime.Add(e.updateFrequency)
		}
	}
}

func (s *sv) CalculationFinished(eventID string, result statevar.StateVariableResult, err error) {
	if err == nil {
		s.result(context.Background(), result)
	}
}

func (e *StateVarStub) ReadyForTimeTrigger(asset, mktID string) {
	if _, ok := e.readyForTimeTrigger[asset+mktID]; !ok {
		e.readyForTimeTrigger[mktID] = struct{}{}
		for _, s := range e.eventTypeToStateVar[statevar.EventTypeTimeTrigger] {
			if s.asset == asset && s.market == mktID {
				e.stateVarToNextCalc[s.ID] = e.currentTime.Add(e.updateFrequency)
			}
		}
	}
}

func (e *StateVarStub) OnTick(ctx context.Context, t time.Time) {
	e.currentTime = t
	stateVarIDs := []string{}
	for ID, nextTime := range e.stateVarToNextCalc {
		if nextTime.UnixNano() <= t.UnixNano() {
			stateVarIDs = append(stateVarIDs, ID)
		}
	}

	sort.Strings(stateVarIDs)
	eventID := t.Format("20060102_150405.999999999")
	for _, ID := range stateVarIDs {
		s := e.svs[ID]
		s.startCalculation(eventID, s)
		e.stateVarToNextCalc[ID] = t.Add(e.updateFrequency)
	}
}

// generate a random 32 chars identifier.
func (e *StateVarStub) generateID(asset, market string) string {
	b := make([]rune, 32)
	for i := range b {
		b[i] = chars[e.rng.Intn(len(chars))]
	}
	return asset + "_" + market + "_" + string(b)
}
