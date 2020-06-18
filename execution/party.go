package execution

import (
	"context"
	"errors"
	"sync"

	"code.vegaprotocol.io/vega/events"
	"code.vegaprotocol.io/vega/logging"
	types "code.vegaprotocol.io/vega/proto"
)

var ErrPartyDoesNotExist = errors.New("party does not exist in party engine")
var ErrNotifyPartyIdMissing = errors.New("notify party id is missing")
var ErrInvalidPartyId = errors.New("party id is not valid")

// Collateral ...
//go:generate go run github.com/golang/mock/mockgen -destination mocks/collateral_mock.go -package mocks code.vegaprotocol.io/vega/execution Collateral
type Collateral interface {
	CreatePartyGeneralAccount(ctx context.Context, partyID, asset string) string
	IncrementBalance(ctx context.Context, id string, amount uint64) error
	DecrementBalance(ctx context.Context, id string, amount uint64) error
	GetAccountByID(id string) (*types.Account, error)
	GetPartyTokenAccount(string) (*types.Account, error)
}

// Party holds the list of parties in the system
type Party struct {
	log           *logging.Logger
	collateral    Collateral
	markets       []types.Market
	broker        Broker
	partyByMarket map[string]map[string]struct{}
	mu            sync.Mutex
}

// NewParty instantiates a new party
func NewParty(log *logging.Logger, col Collateral, markets []types.Market, broker Broker) *Party {
	partyByMarket := map[string]map[string]struct{}{}
	for _, v := range markets {
		partyByMarket[v.Id] = map[string]struct{}{}
	}
	return &Party{
		log:           log,
		collateral:    col,
		markets:       markets,
		broker:        broker,
		partyByMarket: partyByMarket,
	}
}

// GetByMarket returns the list of all the parties in a given market.
func (p *Party) GetByMarket(mktID string) []string {
	parties := p.partyByMarket[mktID]
	out := make([]string, 0, len(parties))
	for k := range parties {
		out = append(out, k)
	}
	return out
}

// GetByMarketAndID searches for a party that exists in the system for the given market.
func (p *Party) GetByMarketAndID(marketID, partyID string) (*types.Party, error) {
	if _, ok := p.partyByMarket[marketID][partyID]; ok {
		return &types.Party{Id: partyID}, nil
	}
	return nil, ErrPartyDoesNotExist
}

// Do we still use this other than in tests?
// NotifyTraderAccountWithTopUpAmount will create a new party in the system
// and top-up it general account with the given amount
func (p *Party) NotifyTraderAccountWithTopUpAmount(ctx context.Context, notify *types.NotifyTraderAccount, amount uint64) error {
	return p.notifyTraderAccount(ctx, notify, amount)
}

// Void type represents nothingness, emptiness
type Void struct{}

// MakeGeneralAccounts creates general accounts on every market for the given party id
func (p *Party) MakeGeneralAccounts(ctx context.Context, partyID string) (map[string]Void, error) {
	if len(partyID) <= 0 {
		return nil, ErrInvalidPartyId
	}

	// ignore errors as they can only happen when the party already exists
	p.broker.Send(events.NewPartyEvent(ctx, types.Party{Id: partyID}))

	result := map[string]Void{}

	for _, mkt := range p.markets {
		p.addParty(partyID, mkt.Id)
		asset, err := mkt.GetAsset()
		if err != nil {
			p.log.Error("unable to get market asset", logging.Error(err))
			return nil, err
		}

		// create account
		// @TODO this context needs to come from somewhere...
		generalAccount := p.collateral.CreatePartyGeneralAccount(ctx, partyID, asset)
		if _, exists := result[generalAccount]; !exists {
			result[generalAccount] = Void{}
			if _, err := p.collateral.GetAccountByID(generalAccount); err != nil {
				p.log.Error("unable to locate created general account",
					logging.String("party-id", partyID),
					logging.String("asset", asset),
					logging.Error(err))
				return nil, err
			}
			if p.log.GetLevel() == logging.DebugLevel {
				p.log.Debug("created general account",
					logging.String("asset", asset),
					logging.String("party-id", partyID))
			}
		}
	}
	return result, nil
}

// NotifyTraderAccount will create a new party in the system
// and top-up it general account with the default amount
func (p *Party) NotifyTraderAccount(ctx context.Context, notify *types.NotifyTraderAccount) error {
	if notify == nil {
		return ErrNotifyPartyIdMissing
	}
	if notify.Amount == 0 {
		return p.notifyTraderAccount(ctx, notify, 1000000000) // 10000.00000
	}
	return p.notifyTraderAccount(ctx, notify, notify.Amount)
}

// returns parties from an existing market (if any)
// @TODO: untie parties from the markets
func (p *Party) getParties() map[string]struct{} {
	var result map[string]struct{}
	for _, result = range p.partyByMarket {
		break // select existing market (if any) parties
	}
	if result == nil {
		result = map[string]struct{}{}
	}
	return result
}

func (p *Party) addMarket(market types.Market) {
	p.mu.Lock()
	if _, found := p.partyByMarket[market.Id]; !found {
		p.markets = append(p.markets, market)
		p.partyByMarket[market.Id] = p.getParties()
	}
	p.mu.Unlock()
}

func (p *Party) addParty(ptyID, mktID string) {
	p.partyByMarket[mktID][ptyID] = struct{}{}
}

func (p *Party) creditGeneralAccount(ctx context.Context, accountID string, amount uint64) error {

	if err := p.collateral.IncrementBalance(ctx, accountID, amount); err != nil {
		p.log.Error("unable to top-up general account", logging.Error(err))
		return err
	}
	acc, err := p.collateral.GetAccountByID(accountID)
	if err != nil {
		p.log.Error("unable to get general account",
			logging.String("party-id", accountID),
			logging.Error(err))
		return err
	}
	if p.log.GetLevel() == logging.DebugLevel {
		p.log.Debug("account top-up",
			logging.String("party-id", accountID),
			logging.Uint64("top-up-amount", amount),
			logging.Uint64("new-balance", acc.Balance))
	}
	return nil
}

func (p *Party) notifyTraderAccount(ctx context.Context, notify *types.NotifyTraderAccount, amount uint64) error {
	if notify == nil {
		return ErrNotifyPartyIdMissing
	}

	generalAccs, err := p.MakeGeneralAccounts(ctx, notify.TraderID)
	if err != nil {
		return err
	}
	for acc := range generalAccs {
		if err = p.creditGeneralAccount(ctx, acc, amount); err != nil {
			return err
		}
	}

	tknAcc, err := p.collateral.GetPartyTokenAccount(notify.TraderID)
	if err != nil {
		return err
	}
	if err := p.collateral.IncrementBalance(ctx, tknAcc.Id, notify.Amount); err != nil {
		p.log.Error("unable to top-up token account", logging.Error(err))
		return err
	}

	return nil
}