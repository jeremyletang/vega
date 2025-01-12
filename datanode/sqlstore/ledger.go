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

package sqlstore

import (
	"context"
	"time"

	"code.vegaprotocol.io/vega/datanode/entities"
	"code.vegaprotocol.io/vega/datanode/metrics"
	"github.com/georgysavva/scany/pgxscan"
)

type Ledger struct {
	*ConnectionSource
	batcher ListBatcher[entities.LedgerEntry]
	pending []entities.LedgerEntry
}

func NewLedger(connectionSource *ConnectionSource) *Ledger {
	a := &Ledger{
		ConnectionSource: connectionSource,
		batcher:          NewListBatcher[entities.LedgerEntry]("ledger", entities.LedgerEntryColumns),
	}
	return a
}

func (ls *Ledger) Flush(ctx context.Context) ([]entities.LedgerEntry, error) {
	defer metrics.StartSQLQuery("Ledger", "Flush")()

	// This creates an entry time for the ledger entry that is guaranteed to be unique
	// Block event sequence number cannot be used as multiple ledger entries can be created
	// as the result of a single transfer event.
	for i, le := range ls.pending {
		le.LedgerEntryTime = entities.CreateLedgerEntryTime(le.VegaTime, i)
		ls.batcher.Add(le)
	}

	ls.pending = nil

	return ls.batcher.Flush(ctx, ls.Connection)
}

func (ls *Ledger) Add(le entities.LedgerEntry) error {
	ls.pending = append(ls.pending, le)
	return nil
}

func (ls *Ledger) GetByLedgerEntryTime(ledgerEntryTime time.Time) (entities.LedgerEntry, error) {
	defer metrics.StartSQLQuery("Ledger", "GetByID")()
	le := entities.LedgerEntry{}
	ctx := context.Background()
	err := pgxscan.Get(ctx, ls.Connection, &le,
		`SELECT ledger_entry_time, account_from_id, account_to_id, quantity, tx_hash, vega_time, transfer_time, type
		 FROM ledger WHERE ledger_entry_time =$1`,
		ledgerEntryTime)
	return le, err
}

func (ls *Ledger) GetAll() ([]entities.LedgerEntry, error) {
	defer metrics.StartSQLQuery("Ledger", "GetAll")()
	ctx := context.Background()
	ledgerEntries := []entities.LedgerEntry{}
	err := pgxscan.Select(ctx, ls.Connection, &ledgerEntries, `
		SELECT ledger_entry_time, account_from_id, account_to_id, quantity, tx_hash, vega_time, transfer_time, type
		FROM ledger`)
	return ledgerEntries, err
}
