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

package entities

import (
	"fmt"
	"time"

	"code.vegaprotocol.io/vega/protos/vega"
)

const (
	noMarketStr     string = "!"
	noMarketByte    byte   = '!'
	systemOwnerStr  string = "*"
	systemOwnerByte byte   = '*'
)

type (
	_Account  struct{}
	AccountID = ID[_Account]
)

type Account struct {
	ID       AccountID
	PartyID  PartyID
	AssetID  AssetID
	MarketID MarketID
	Type     vega.AccountType
	TxHash   TxHash
	VegaTime time.Time
}

func (a Account) String() string {
	return fmt.Sprintf("{ID: %s}", a.ID)
}

func AccountFromProto(va *vega.Account, txHash TxHash) (Account, error) {
	account := Account{
		PartyID:  PartyID(va.Owner),
		AssetID:  AssetID(va.Asset),
		MarketID: MarketID(va.MarketId),
		Type:     va.Type,
		TxHash:   txHash,
	}
	return account, nil
}

func AccountProtoFromDetails(ad *vega.AccountDetails, txHash TxHash) (Account, error) {
	marketID, partyID := noMarketStr, systemOwnerStr
	if ad.MarketId != nil {
		marketID = *ad.MarketId
	}
	if ad.Owner != nil {
		partyID = *ad.Owner
	}
	return Account{
		TxHash:   txHash,
		PartyID:  ID[_Party](partyID),
		MarketID: ID[_Market](marketID),
		Type:     ad.Type,
		AssetID:  ID[_Asset](ad.AssetId),
	}, nil
}
