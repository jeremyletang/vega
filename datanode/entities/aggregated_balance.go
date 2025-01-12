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
	"encoding/json"
	"fmt"
	"time"

	"code.vegaprotocol.io/vega/core/types"
	v2 "code.vegaprotocol.io/vega/protos/data-node/api/v2"
	"code.vegaprotocol.io/vega/protos/vega"
	"github.com/shopspring/decimal"
)

// AggregatedBalance represents the the summed balance of a bunch of accounts at a given
// time. VegaTime and Balance will always be set. The others will be nil unless when
// querying, you requested grouping by one of the corresponding fields.
type AggregatedBalance struct {
	VegaTime  time.Time
	Balance   decimal.Decimal
	AccountID AccountID
	PartyID   PartyID
	AssetID   AssetID
	MarketID  MarketID
	Type      *types.AccountType
}

func (balance *AggregatedBalance) ToProto() *v2.AggregatedBalance {
	var accountType vega.AccountType
	var accountID, partyID, assetID, marketID *string

	if balance.AccountID != "" {
		aid := balance.AccountID.String()
		accountID = &aid
	}

	if balance.PartyID != "" {
		pid := balance.PartyID.String()
		partyID = &pid
	}

	if balance.AssetID != "" {
		aid := balance.AssetID.String()
		assetID = &aid
	}

	if balance.MarketID != "" {
		mid := balance.MarketID.String()
		marketID = &mid
	}

	if balance.Type != nil {
		accountType = *balance.Type
	}

	return &v2.AggregatedBalance{
		Timestamp:   balance.VegaTime.UnixNano(),
		Balance:     balance.Balance.String(),
		AccountId:   accountID,
		PartyId:     partyID,
		AssetId:     assetID,
		MarketId:    marketID,
		AccountType: accountType,
	}
}

func (balance AggregatedBalance) Cursor() *Cursor {
	return NewCursor(AggregatedBalanceCursor{
		VegaTime: balance.VegaTime,
	}.String())
}

func (balance AggregatedBalance) ToProtoEdge(_ ...any) (*v2.AggregatedBalanceEdge, error) {
	return &v2.AggregatedBalanceEdge{
		Node:   balance.ToProto(),
		Cursor: balance.Cursor().Encode(),
	}, nil
}

type AggregatedBalanceCursor struct {
	VegaTime time.Time `json:"vega_time"`
}

func (c AggregatedBalanceCursor) String() string {
	bs, err := json.Marshal(c)
	if err != nil {
		panic(fmt.Errorf("could not marshal aggregate balance cursor: %w", err))
	}
	return string(bs)
}

func (c *AggregatedBalanceCursor) Parse(cursorString string) error {
	if cursorString == "" {
		return nil
	}
	return json.Unmarshal([]byte(cursorString), c)
}
