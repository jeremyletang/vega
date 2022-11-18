// Copyright (c) 2022 Gobalsky Labs Limited
//
// Use of this software is governed by the Business Source License included
// in the LICENSE.DATANODE file and at https://www.mariadb.com/bsl11.
//
// Change Date: 18 months from the later of the date of the first publicly
// available Distribution of this version of the repository, and 25 June 2022.
//
// On the date above, in accordance with the Business Source License, use
// of this software will be governed by version 3 or later of the GNU General
// Public License.

package service

import (
	"context"

	"code.vegaprotocol.io/vega/datanode/entities"
	"code.vegaprotocol.io/vega/datanode/utils"
	"code.vegaprotocol.io/vega/logging"
	"code.vegaprotocol.io/vega/protos/vega"
)

type ledgerStore interface {
	Flush(ctx context.Context) ([]entities.LedgerEntry, error)
	Add(le entities.LedgerEntry) error
	Query(filter *entities.LedgerEntryFilter, dateRange entities.DateRange, pagination entities.CursorPagination) (*[]entities.AggregatedLedgerEntries, entities.PageInfo, error)
}

type LedgerEntriesStore interface {
	Query(filter *entities.LedgerEntryFilter, dateRange entities.DateRange, pagination entities.CursorPagination) (*[]entities.AggregatedLedgerEntries, entities.PageInfo, error)
}

type Ledger struct {
	store             ledgerStore
	log               *logging.Logger
	transferResponses []*vega.LedgerMovement
	observer          utils.Observer[*vega.LedgerMovement]
}

func NewLedger(store ledgerStore, log *logging.Logger) *Ledger {
	return &Ledger{
		store:    store,
		log:      log,
		observer: utils.NewObserver[*vega.LedgerMovement]("ledger", log, 0, 0),
	}
}

func (l *Ledger) Flush(ctx context.Context) error {
	_, err := l.store.Flush(ctx)
	if err != nil {
		return err
	}
	l.observer.Notify(l.transferResponses)
	l.transferResponses = []*vega.LedgerMovement{}
	return nil
}

func (l *Ledger) AddLedgerEntry(le entities.LedgerEntry) error {
	return l.store.Add(le)
}

func (l *Ledger) AddTransferResponse(le *vega.LedgerMovement) {
	l.transferResponses = append(l.transferResponses, le)
}

func (l *Ledger) Observe(ctx context.Context, retries int) (<-chan []*vega.LedgerMovement, uint64) {
	ch, ref := l.observer.Observe(ctx,
		retries,
		func(tr *vega.LedgerMovement) bool {
			return true
		})
	return ch, ref
}

func (l *Ledger) GetSubscribersCount() int32 {
	return l.observer.GetSubscribersCount()
}

func (l *Ledger) Query(
	filter *entities.LedgerEntryFilter,
	dateRange entities.DateRange,
	pagination entities.CursorPagination,
) (*[]entities.AggregatedLedgerEntries, entities.PageInfo, error) {
	return l.store.Query(
		filter,
		dateRange,
		pagination)
}
