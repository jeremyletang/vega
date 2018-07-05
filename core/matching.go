package core

import (
	"vega/matching"
	"vega/proto"
)

type MatchingEngine interface {
	CreateMarket(id string)
	SubmitOrder(order msg.Order) (*msg.OrderConfirmation, msg.OrderError)
	DeleteOrder(id string) *msg.Order
}

func (v Vega) CreateMarket(id string) {
	if _, exists := v.markets[id]; !exists {
		book := matching.NewBook(id, v.config.Matching)
		v.markets[id] = book
	}
}

func (v Vega) SubmitOrder(order *msg.Order) (*msg.OrderConfirmation, msg.OrderError) {
	market, exists := v.markets[order.Market]
	if !exists {
		return nil, msg.OrderError_INVALID_MARKET_ID
	}

	confirmationMessage, err := market.AddOrder(order)
	if err != msg.OrderError_NONE {
		return nil, err
	}
	return confirmationMessage, msg.OrderError_NONE
}

func (v Vega) DeleteOrder(order *msg.Order) {
	if market, exists := v.markets[order.Market]; exists {
		market.RemoveOrder(order)
	}
}
