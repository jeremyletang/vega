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

package sqlstore_test

import (
	"context"
	"testing"
	"time"

	"code.vegaprotocol.io/vega/datanode/entities"
	"code.vegaprotocol.io/vega/datanode/sqlstore"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func addTestLedgerEntry(t *testing.T, ledger *sqlstore.Ledger,
	accountFrom entities.Account,
	accountTo entities.Account,
	block entities.Block,
) entities.LedgerEntry {
	t.Helper()
	ledgerEntry := entities.LedgerEntry{
		AccountFromID: accountFrom.ID,
		AccountToID:   accountTo.ID,
		Quantity:      decimal.NewFromInt(100),
		VegaTime:      block.VegaTime,
		TransferTime:  block.VegaTime.Add(-time.Second),
		Type:          entities.LedgerMovementTypeBondSlashing,
	}

	err := ledger.Add(ledgerEntry)
	require.NoError(t, err)
	return ledgerEntry
}

func TestLedger(t *testing.T) {
	defer DeleteEverything()
	ctx := context.Background()

	blockStore := sqlstore.NewBlocks(connectionSource)
	assetStore := sqlstore.NewAssets(connectionSource)
	accountStore := sqlstore.NewAccounts(connectionSource)
	partyStore := sqlstore.NewParties(connectionSource)
	ledgerStore := sqlstore.NewLedger(connectionSource)

	// Account store should be empty to begin with
	ledgerEntries, err := ledgerStore.GetAll()
	assert.NoError(t, err)
	assert.Empty(t, ledgerEntries)

	block := addTestBlock(t, blockStore)
	asset := addTestAsset(t, assetStore, block)
	party := addTestParty(t, partyStore, block)
	accountFrom := addTestAccount(t, accountStore, party, asset, block)
	accountTo := addTestAccount(t, accountStore, party, asset, block)
	firstLedgerEntry := addTestLedgerEntry(t, ledgerStore, accountFrom, accountTo, block)

	_, err = ledgerStore.Flush(ctx)
	assert.NoError(t, err)

	// Add it again twice; we're allowed multiple ledger entries with the same parameters in the same block as well
	// as acroos block
	block2 := addTestBlock(t, blockStore)
	addTestLedgerEntry(t, ledgerStore, accountFrom, accountTo, block2)
	addTestLedgerEntry(t, ledgerStore, accountFrom, accountTo, block2)

	_, err = ledgerStore.Flush(ctx)
	assert.NoError(t, err)

	// Query and check we've got back an asset the same as the one we put in, once we give it an ID
	ledgerEntryTime := entities.CreateLedgerEntryTime(block.VegaTime, 0)

	fetchedLedgerEntry, err := ledgerStore.GetByLedgerEntryTime(ledgerEntryTime)
	assert.NoError(t, err)
	firstLedgerEntry.LedgerEntryTime = ledgerEntryTime
	assert.Equal(t, firstLedgerEntry, fetchedLedgerEntry)

	// We should have added three entries in total
	ledgerEntriesAfter, err := ledgerStore.GetAll()
	assert.NoError(t, err)
	assert.Len(t, ledgerEntriesAfter, 3)
}
