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

	v2 "code.vegaprotocol.io/vega/protos/data-node/api/v2"

	"code.vegaprotocol.io/vega/core/events"
	"code.vegaprotocol.io/vega/libs/num"
	"code.vegaprotocol.io/vega/protos/vega"
	"github.com/shopspring/decimal"
)

type positionSettlement interface {
	Price() *num.Uint
	PositionFactor() num.Decimal
	Trades() []events.TradeSettlement
	TxHash() string
}

type lossSocialization interface {
	Amount() *num.Int
	TxHash() string
}

type settleDistressed interface {
	Margin() *num.Uint
	TxHash() string
}

type settleMarket interface {
	SettledPrice() *num.Uint
	PositionFactor() num.Decimal
	TxHash() string
}

type Position struct {
	MarketID                MarketID
	PartyID                 PartyID
	OpenVolume              int64
	RealisedPnl             decimal.Decimal
	UnrealisedPnl           decimal.Decimal
	AverageEntryPrice       decimal.Decimal
	AverageEntryMarketPrice decimal.Decimal
	Loss                    decimal.Decimal // what the party lost because of loss socialization
	Adjustment              decimal.Decimal // what a party was missing which triggered loss socialization
	TxHash                  TxHash
	VegaTime                time.Time
}

func NewEmptyPosition(marketID MarketID, partyID PartyID) Position {
	return Position{
		MarketID:                marketID,
		PartyID:                 partyID,
		OpenVolume:              0,
		RealisedPnl:             num.DecimalZero(),
		UnrealisedPnl:           num.DecimalZero(),
		AverageEntryPrice:       num.DecimalZero(),
		AverageEntryMarketPrice: num.DecimalZero(),
		Loss:                    num.DecimalZero(),
		Adjustment:              num.DecimalZero(),
	}
}

func (p *Position) UpdateWithPositionSettlement(e positionSettlement) {
	for _, t := range e.Trades() {
		openedVolume, closedVolume := calculateOpenClosedVolume(p.OpenVolume, t.Size())
		// Deal with any volume we have closed
		realisedPnlDelta := num.DecimalFromUint(t.Price()).Sub(p.AverageEntryPrice).Mul(num.DecimalFromInt64(closedVolume)).Div(e.PositionFactor())
		p.RealisedPnl = p.RealisedPnl.Add(realisedPnlDelta)
		p.OpenVolume -= closedVolume

		// Then with any we have opened
		p.AverageEntryPrice = updateVWAP(p.AverageEntryPrice, p.OpenVolume, openedVolume, t.Price())
		p.AverageEntryMarketPrice = updateVWAP(p.AverageEntryMarketPrice, p.OpenVolume, openedVolume, t.MarketPrice())
		p.OpenVolume += openedVolume
	}
	p.mtm(e.Price(), e.PositionFactor())
	p.TxHash = TxHash(e.TxHash())
}

func (p *Position) UpdateWithLossSocialization(e lossSocialization) {
	amountLoss := num.DecimalFromInt(e.Amount())

	if amountLoss.IsNegative() {
		p.Loss = p.Loss.Add(amountLoss)
	} else {
		p.Adjustment = p.Adjustment.Add(amountLoss)
	}

	p.RealisedPnl = p.RealisedPnl.Add(amountLoss)
	p.TxHash = TxHash(e.TxHash())
}

func (p *Position) UpdateWithSettleDistressed(e settleDistressed) {
	margin := num.DecimalFromUint(e.Margin())
	p.RealisedPnl = p.RealisedPnl.Add(p.UnrealisedPnl)
	p.RealisedPnl = p.RealisedPnl.Sub(margin) // realised P&L includes whatever we had in margin account at this point
	p.UnrealisedPnl = num.DecimalZero()
	p.AverageEntryPrice = num.DecimalZero() // @TODO average entry price shouldn't be affected(?)
	p.AverageEntryPrice = num.DecimalZero()
	p.OpenVolume = 0
	p.TxHash = TxHash(e.TxHash())
}

func (p *Position) UpdateWithSettleMarket(e settleMarket) {
	markPriceDec := num.DecimalFromUint(e.SettledPrice())
	openVolumeDec := num.DecimalFromInt64(p.OpenVolume)

	unrealisedPnl := openVolumeDec.Mul(markPriceDec.Sub(p.AverageEntryPrice)).Div(e.PositionFactor())
	p.RealisedPnl = p.RealisedPnl.Add(unrealisedPnl)
	p.UnrealisedPnl = num.DecimalZero()
	p.OpenVolume = 0
	p.TxHash = TxHash(e.TxHash())
}

func (p *Position) ToProto() *vega.Position {
	var timestamp int64
	if !p.VegaTime.IsZero() {
		timestamp = p.VegaTime.UnixNano()
	}
	return &vega.Position{
		MarketId:          p.MarketID.String(),
		PartyId:           p.PartyID.String(),
		OpenVolume:        p.OpenVolume,
		RealisedPnl:       p.RealisedPnl.Round(0).String(),
		UnrealisedPnl:     p.UnrealisedPnl.Round(0).String(),
		AverageEntryPrice: p.AverageEntryMarketPrice.Round(0).String(),
		UpdatedAt:         timestamp,
	}
}

func (p Position) ToProtoEdge(_ ...any) (*v2.PositionEdge, error) {
	return &v2.PositionEdge{
		Node:   p.ToProto(),
		Cursor: p.Cursor().Encode(),
	}, nil
}

func (p *Position) AverageEntryPriceUint() *num.Uint {
	uint, overflow := num.UintFromDecimal(p.AverageEntryPrice)
	if overflow {
		panic("couldn't convert average entry price from decimal to uint")
	}
	return uint
}

func (p *Position) mtm(markPrice *num.Uint, positionFactor num.Decimal) {
	if p.OpenVolume == 0 {
		p.UnrealisedPnl = num.DecimalZero()
		return
	}
	markPriceDec := num.DecimalFromUint(markPrice)
	openVolumeDec := num.DecimalFromInt64(p.OpenVolume)

	p.UnrealisedPnl = openVolumeDec.Mul(markPriceDec.Sub(p.AverageEntryPrice)).Div(positionFactor)
}

func calculateOpenClosedVolume(currentOpenVolume, tradedVolume int64) (int64, int64) {
	if currentOpenVolume != 0 && ((currentOpenVolume > 0) != (tradedVolume > 0)) {
		var closedVolume int64
		if absUint64(tradedVolume) > absUint64(currentOpenVolume) {
			closedVolume = currentOpenVolume
		} else {
			closedVolume = -tradedVolume
		}
		return tradedVolume + closedVolume, closedVolume
	}
	return tradedVolume, 0
}

func absUint64(v int64) uint64 {
	if v < 0 {
		v *= -1
	}
	return uint64(v)
}

func updateVWAP(vwap num.Decimal, volume int64, addVolume int64, addPrice *num.Uint) num.Decimal {
	if volume+addVolume == 0 {
		return num.DecimalZero()
	}

	volumeDec := num.DecimalFromInt64(volume)
	addVolumeDec := num.DecimalFromInt64(addVolume)
	addPriceDec := num.DecimalFromUint(addPrice)

	return vwap.Mul(volumeDec).Add(addPriceDec.Mul(addVolumeDec)).Div(volumeDec.Add(addVolumeDec))
}

type PositionKey struct {
	MarketID MarketID
	PartyID  PartyID
	VegaTime time.Time
}

func (p Position) Cursor() *Cursor {
	pc := PositionCursor{
		MarketID: p.MarketID,
		PartyID:  p.PartyID,
		VegaTime: p.VegaTime,
	}

	return NewCursor(pc.String())
}

func (p Position) Key() PositionKey {
	return PositionKey{p.MarketID, p.PartyID, p.VegaTime}
}

var PositionColumns = []string{
	"market_id", "party_id", "open_volume", "realised_pnl", "unrealised_pnl",
	"average_entry_price", "average_entry_market_price", "loss", "adjustment", "tx_hash", "vega_time",
}

func (p Position) ToRow() []interface{} {
	return []interface{}{
		p.MarketID, p.PartyID, p.OpenVolume, p.RealisedPnl, p.UnrealisedPnl,
		p.AverageEntryPrice, p.AverageEntryMarketPrice, p.Loss, p.Adjustment, p.TxHash, p.VegaTime,
	}
}

func (p Position) Equal(q Position) bool {
	return p.MarketID == q.MarketID &&
		p.PartyID == q.PartyID &&
		p.OpenVolume == q.OpenVolume &&
		p.RealisedPnl.Equal(q.RealisedPnl) &&
		p.UnrealisedPnl.Equal(q.UnrealisedPnl) &&
		p.AverageEntryPrice.Equal(q.AverageEntryPrice) &&
		p.AverageEntryMarketPrice.Equal(q.AverageEntryMarketPrice) &&
		p.Loss.Equal(q.Loss) &&
		p.Adjustment.Equal(q.Adjustment) &&
		p.TxHash == q.TxHash &&
		p.VegaTime.Equal(q.VegaTime)
}

type PositionCursor struct {
	VegaTime time.Time `json:"vega_time"`
	PartyID  PartyID   `json:"party_id"`
	MarketID MarketID  `json:"market_id"`
}

func (rc PositionCursor) String() string {
	bs, err := json.Marshal(rc)
	if err != nil {
		// This should never happen.
		panic(fmt.Errorf("could not marshal order cursor: %w", err))
	}
	return string(bs)
}

func (rc *PositionCursor) Parse(cursorString string) error {
	if cursorString == "" {
		return nil
	}
	return json.Unmarshal([]byte(cursorString), rc)
}
