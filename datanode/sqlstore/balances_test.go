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
	"fmt"
	"testing"
	"time"

	"code.vegaprotocol.io/vega/datanode/entities"
	"code.vegaprotocol.io/vega/datanode/sqlstore"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func addTestBalance(t *testing.T, store *sqlstore.Balances, block entities.Block, acc entities.Account, balance int64) {
	t.Helper()
	bal := entities.AccountBalance{
		Account:  &acc,
		VegaTime: block.VegaTime,
		Balance:  decimal.NewFromInt(balance),
	}

	err := store.Add(bal)
	require.NoError(t, err)
}

func assertBalanceCorrect(t *testing.T,
	expectedBlocks []int, expectedBals []int64,
	blocks []entities.Block, bals []entities.AggregatedBalance,
) {
	t.Helper()
	assert.Len(t, bals, len(expectedBlocks))
	for i := 0; i < len(expectedBlocks); i++ {
		assert.Equal(t, blocks[expectedBlocks[i]].VegaTime, (bals)[i].VegaTime)
		assert.Equal(t, decimal.NewFromInt(expectedBals[i]), (bals)[i].Balance)
	}
}

func TestBalances(t *testing.T) {
	defer DeleteEverything()
	ctx := context.Background()

	blockStore := sqlstore.NewBlocks(connectionSource)
	assetStore := sqlstore.NewAssets(connectionSource)
	accountStore := sqlstore.NewAccounts(connectionSource)
	balanceStore := sqlstore.NewBalances(connectionSource)
	partyStore := sqlstore.NewParties(connectionSource)

	// Set up a test environment with a bunch of blocks/parties/accounts
	asset := addTestAsset(t, assetStore, addTestBlock(t, blockStore))

	var blocks []entities.Block
	var parties []entities.Party
	var accounts []entities.Account
	for i := 0; i < 5; i++ {
		blocks = append(blocks, addTestBlock(t, blockStore))
		parties = append(parties, addTestParty(t, partyStore, blocks[0]))
		accounts = append(accounts, addTestAccount(t, accountStore, parties[i], asset, blocks[0]))
	}

	// And add some dummy balances
	addTestBalance(t, balanceStore, blocks[0], accounts[0], 1)
	addTestBalance(t, balanceStore, blocks[0], accounts[0], 2) // Second balance on same acc/block should override first
	addTestBalance(t, balanceStore, blocks[1], accounts[0], 5)
	addTestBalance(t, balanceStore, blocks[2], accounts[1], 10)
	addTestBalance(t, balanceStore, blocks[3], accounts[2], 100)
	addTestBalance(t, balanceStore, blocks[4], accounts[0], 30)

	balanceStore.Flush(ctx)

	dateRange := entities.DateRange{}
	pagination := entities.CursorPagination{}

	t.Run("Query should return all balances", func(t *testing.T) {
		// Query all the balances (they're all for the same asset)
		bals, _, err := balanceStore.Query(entities.AccountFilter{AssetID: asset.ID}, []entities.AccountField{}, dateRange, pagination)
		require.NoError(t, err)

		expectedBlocks := []int{0, 1, 2, 3, 4}
		expectedBals := []int64{2, 5, 5 + 10, 5 + 10 + 100, 30 + 10 + 100}
		assertBalanceCorrect(t, expectedBlocks, expectedBals, blocks[:], *bals)
	})
	/* TODO Phil is going to refactor these tests to not depend on the account id order
		t.Run("Query should return transactions for party", func(t *testing.T) {
			// Try just for our first account/party
			filter := entities.AccountFilter{
				AssetID:  asset.ID,
				PartyIDs: []entities.PartyID{parties[0].ID},
			}
			bals, _, err := balanceStore.Query(filter, []entities.AccountField{}, dateRange, pagination)
			require.NoError(t, err)

			expectedBlocks := []int{0, 1, 4}
			expectedBals := []int64{2, 5, 30}
			assertBalanceCorrect(t, expectedBlocks, expectedBals, blocks[:], *bals)
		})


	t.Run("Query should group results", func(t *testing.T) {
		// Now try grouping - if we do it by account id it should split out balances for each account.
		bals, _, err := balanceStore.Query(entities.AccountFilter{AssetID: asset.ID}, []entities.AccountField{entities.AccountFieldID}, dateRange, pagination)
		require.NoError(t, err)

		expectedBlocks := []int{0, 1, 2, 2, 3, 3, 3, 4, 4, 4}
		expectedBals := []int64{2, 5, 5, 10, 5, 10, 100, 30, 10, 100}
		assertBalanceCorrect(t, expectedBlocks, expectedBals, blocks[:], *bals)
	})


		t.Run("Query should return results paged", func(t *testing.T) {
			first := int32(3)
			after := entities.NewCursor(entities.AggregatedBalanceCursor{
				VegaTime: blocks[2].VegaTime,
			}.String()).Encode()
			p, err := entities.NewCursorPagination(&first, &after, nil, nil, false)
			require.NoError(t, err)
			bals, _, err := balanceStore.Query(entities.AccountFilter{AssetID: asset.ID}, []entities.AccountField{entities.AccountFieldID}, dateRange, p)
			require.NoError(t, err)

			expectedBlocks := []int{2, 3, 3}
			expectedBals := []int64{10, 5, 10}
			assertBalanceCorrect(t, expectedBlocks, expectedBals, blocks[:], *bals)
		})

		t.Run("Query should return results between dates", func(t *testing.T) {
			p, err := entities.NewCursorPagination(nil, nil, nil, nil, false)
			require.NoError(t, err)
			startTime := blocks[1].VegaTime
			endTime := blocks[4].VegaTime
			dateRange := entities.DateRange{
				Start: &startTime,
				End:   &endTime,
			}
			bals, _, err := balanceStore.Query(entities.AccountFilter{AssetID: asset.ID}, []entities.AccountField{entities.AccountFieldID}, dateRange, p)
			require.NoError(t, err)

			expectedBlocks := []int{1, 2, 2, 3, 3, 3}
			expectedBals := []int64{5, 5, 10, 5, 10, 100}
			assertBalanceCorrect(t, expectedBlocks, expectedBals, blocks[:], *bals)
		})

		t.Run("Query should return results paged between dates", func(t *testing.T) {
			first := int32(3)
			p, err := entities.NewCursorPagination(&first, nil, nil, nil, false)
			require.NoError(t, err)
			startTime := blocks[1].VegaTime
			endTime := blocks[4].VegaTime
			dateRange := entities.DateRange{
				Start: &startTime,
				End:   &endTime,
			}
			bals, _, err := balanceStore.Query(entities.AccountFilter{AssetID: asset.ID}, []entities.AccountField{entities.AccountFieldID}, dateRange, p)
			require.NoError(t, err)

			expectedBlocks := []int{1, 2, 2}
			expectedBals := []int64{5, 5, 10}
			assertBalanceCorrect(t, expectedBlocks, expectedBals, blocks[:], *bals)
		})
	*/
}

func TestBalancesDataRetention(t *testing.T) {
	defer DeleteEverything()
	ctx := context.Background()

	blockStore := sqlstore.NewBlocks(connectionSource)
	assetStore := sqlstore.NewAssets(connectionSource)
	accountStore := sqlstore.NewAccounts(connectionSource)
	balanceStore := sqlstore.NewBalances(connectionSource)
	partyStore := sqlstore.NewParties(connectionSource)

	// Set up a test environment with a bunch of blocks/parties/accounts
	asset := addTestAsset(t, assetStore, addTestBlock(t, blockStore))

	var blocks []entities.Block
	var parties []entities.Party
	var accounts []entities.Account
	for i := 0; i < 5; i++ {
		blocks = append(blocks, addTestBlockForTime(t, blockStore, time.Now().Add((-26*time.Hour)-(time.Duration(5-i)*time.Second))))
		parties = append(parties, addTestParty(t, partyStore, blocks[0]))
		accounts = append(accounts, addTestAccount(t, accountStore, parties[i], asset, blocks[0]))
	}

	// And add some dummy balances
	addTestBalance(t, balanceStore, blocks[0], accounts[0], 1)
	addTestBalance(t, balanceStore, blocks[0], accounts[0], 2) // Second balance on same acc/block should override first
	addTestBalance(t, balanceStore, blocks[1], accounts[0], 5)
	addTestBalance(t, balanceStore, blocks[2], accounts[1], 10)
	addTestBalance(t, balanceStore, blocks[3], accounts[2], 100)
	balanceStore.Flush(ctx)

	// Conflate the data and add some new positions so all tests run against a mix of conflated and non-conflated data
	now := time.Now()
	refreshQuery := fmt.Sprintf("CALL refresh_continuous_aggregate('conflated_balances', '%s', '%s');",
		now.Add(-48*time.Hour).Format("2006-01-02"),
		time.Now().Format("2006-01-02"))
	_, err := connectionSource.Connection.Exec(context.Background(), refreshQuery)

	assert.NoError(t, err)

	// The refresh of the continuous aggregate completes asynchronously so the following loop is necessary to ensure the data has been materialized
	// before the test continues
	for {
		var counter int
		connectionSource.Connection.QueryRow(context.Background(), "SELECT count(*) FROM conflated_balances").Scan(&counter)
		if counter == 3 {
			break
		}
	}

	_, err = connectionSource.Connection.Exec(context.Background(), "delete from balances")
	assert.NoError(t, err)

	addTestBalance(t, balanceStore, blocks[4], accounts[0], 30)
	balanceStore.Flush(ctx)

	dateRange := entities.DateRange{}
	pagination := entities.CursorPagination{}

	// Query all the balances (they're all for the same asset)
	bals, _, err := balanceStore.Query(entities.AccountFilter{AssetID: asset.ID}, []entities.AccountField{}, dateRange, pagination)
	require.NoError(t, err)

	expectedBlocks := []int{1, 2, 3, 4}
	expectedBals := []int64{5, 5 + 10, 5 + 10 + 100, 30 + 10 + 100}
	assertBalanceCorrect(t, expectedBlocks, expectedBals, blocks[:], *bals)
}
