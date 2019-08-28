package storage_test

import (
	types "code.vegaprotocol.io/vega/proto"
	"testing"

	"code.vegaprotocol.io/vega/internal/config/encoding"
	"code.vegaprotocol.io/vega/internal/logging"
	"code.vegaprotocol.io/vega/internal/storage"

	"github.com/stretchr/testify/assert"
)


const (
	testAccountParty1  = "g0ldman"
	testAccountParty2  = "m3tr0"
	testAccountStore   = "accountstore-test"
	testAccountMarket1 = "m@rk3t1"
	testAccountMarket2 = "tr@d1nG"
	testAssetGBP       = "GBP"
	testAssetUSD       = "USD"
	testAssetEUR       = "EUR"
)

func TestAccount_GetByPartyAndAsset(t *testing.T) {
	dir, tidy := createTmpDir(t, testAccountStore)
	defer tidy()

	accountStore := createAccountStore(t, dir)

	err := accountStore.SaveBatch(getTestAccounts())
	assert.Nil(t, err)

	accs, err := accountStore.GetByPartyAndAsset(testAccountParty2, testAssetEUR)
	assert.Nil(t, err)
	assert.Len(t, accs, 2)
	assert.Equal(t, accs[0].Asset, testAssetEUR)
	assert.Equal(t, accs[1].Asset, testAssetEUR)

	accs, err = accountStore.GetByPartyAndAsset(testAccountParty1, testAssetEUR)
	assert.Nil(t, err)
	assert.Len(t, accs, 0)

	accs, err = accountStore.GetByPartyAndAsset(testAccountParty1, testAssetUSD)
	assert.Nil(t, err)
	assert.Len(t, accs, 2)
	assert.Equal(t, accs[0].Asset, testAssetUSD)
	assert.Equal(t, accs[1].Asset, testAssetUSD)

	accs, err = accountStore.GetByPartyAndAsset(testAccountParty1, testAssetGBP)
	assert.Nil(t, err)
	assert.Len(t, accs, 2)
	assert.Equal(t, accs[0].Asset, testAssetGBP)
	assert.Equal(t, accs[1].Asset, testAssetGBP)

	err = accountStore.Close()
	assert.NoError(t, err)
}

func TestAccount_GetByPartyAndType(t *testing.T) {
	invalid := "invalid type for query"

	dir, tidy := createTmpDir(t, testAccountStore)
	defer tidy()

	accountStore := createAccountStore(t, dir)

	err := accountStore.SaveBatch(getTestAccounts())
	assert.Nil(t, err)

	// General accounts
	accs, err := accountStore.GetByPartyAndType(testAccountParty1, types.AccountType_GENERAL)
	assert.Nil(t, err)
	assert.Len(t, accs, 2)

	// Margin accounts
	accs, err = accountStore.GetByPartyAndType(testAccountParty1, types.AccountType_MARGIN)
	assert.Nil(t, err)
	assert.Len(t, accs, 2)

	// Invalid account type
	accs, err = accountStore.GetByPartyAndType(testAccountParty2, types.AccountType_INSURANCE)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), invalid)
	assert.Nil(t, accs)

	// Invalid account type
	accs, err = accountStore.GetByPartyAndType(testAccountParty2, types.AccountType_SETTLEMENT)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), invalid)
	assert.Nil(t, accs)

	err = accountStore.Close()
	assert.NoError(t, err)
}

func TestAccount_GetByPartyAndMarket(t *testing.T) {
	dir, tidy := createTmpDir(t, testAccountStore)
	defer tidy()

	accountStore := createAccountStore(t, dir)

	err := accountStore.SaveBatch(getTestAccounts())
	assert.Nil(t, err)

	accs, err := accountStore.GetByPartyAndMarket(testAccountParty1, testAccountMarket1)
	assert.Nil(t, err)
	assert.Len(t, accs, 1)
	assert.Equal(t, testAccountMarket1, accs[0].MarketID)

	accs, err = accountStore.GetByPartyAndMarket(testAccountParty1, testAccountMarket2)
	assert.Nil(t, err)
	assert.Len(t, accs, 1)
	assert.Equal(t, testAccountMarket2, accs[0].MarketID)

	err = accountStore.Close()
	assert.NoError(t, err)
}

func TestAccount_GetByParty(t *testing.T) {
	dir, tidy := createTmpDir(t, testAccountStore)
	defer tidy()

	accountStore := createAccountStore(t, dir)

	err := accountStore.SaveBatch(getTestAccounts())
	assert.Nil(t, err)

	accs, err := accountStore.GetByParty(testAccountParty1)
	assert.Nil(t, err)
	assert.Len(t, accs, 4)
	assert.Equal(t, testAccountMarket1, accs[0].MarketID)

	err = accountStore.Close()
	assert.NoError(t, err)
}

func getTestAccounts() []*types.Account {
	accs := []*types.Account {
		{
			Owner: testAccountParty1,
			MarketID: testAccountMarket1,
			Type: types.AccountType_GENERAL,
			Asset: testAssetGBP,
			Balance: 1024,
		},
		{
			Owner: testAccountParty1,
			MarketID: testAccountMarket1,
			Type: types.AccountType_MARGIN,
			Asset: testAssetGBP,
			Balance: 1024,
		},
		{
			Owner: testAccountParty1,
			MarketID: testAccountMarket2,
			Type: types.AccountType_GENERAL,
			Asset: testAssetUSD,
			Balance: 1,
		},
		{
			Owner: testAccountParty1,
			MarketID: testAccountMarket2,
			Type: types.AccountType_MARGIN,
			Asset: testAssetUSD,
			Balance: 9,
		},
		{
			Owner: testAccountParty2,
			MarketID: testAccountMarket2,
			Type: types.AccountType_GENERAL,
			Asset: testAssetEUR,
			Balance: 2048,
		},
		{
			Owner: testAccountParty2,
			MarketID: testAccountMarket2,
			Type: types.AccountType_MARGIN,
			Asset: testAssetEUR,
			Balance: 1024,
		},
	}
	return accs
}

func createTmpDir(t *testing.T, storePath string) (string, func()) {
	dir, tidy, err := storage.TempDir(storePath)
	if err != nil {
		t.Fatalf("Failed to create tmp dir: %s", err.Error())
	}
	return dir, tidy
}

func createAccountStore(t *testing.T, dir string) *storage.Account {
	config := storage.Config{
		Level:           encoding.LogLevel{Level: logging.DebugLevel},
		Accounts:        storage.DefaultAccountStoreOptions(),
		AccountsDirPath: dir,
	}
	accountStore, err := storage.NewAccounts(logging.NewTestLogger(), config)
	assert.NoError(t, err)
	assert.NotNil(t, accountStore)

	if accountStore == nil {
		t.Fatalf("Error creating account store in unit test(s)")
	}

	return accountStore
}
