package execution

import (
	"errors"
	"sync"

	"code.vegaprotocol.io/vega/logging"
	"code.vegaprotocol.io/vega/proto"
	types "code.vegaprotocol.io/vega/proto"
)

var ErrPartyDoesNotExist = errors.New("party does not exist in party engine")

// Collateral ...
//go:generate go run github.com/golang/mock/mockgen -destination mocks/collateral_mock.go -package mocks code.vegaprotocol.io/vega/execution Collateral
type Collateral interface {
	CreatePartyGeneralAccount(partyID, asset string) string
	IncrementBalance(id string, amount int64) error
	GetAccountByID(id string) (*proto.Account, error)
}

// Party holds the list of parties in the system
type Party struct {
	log           *logging.Logger
	collateral    Collateral
	markets       []proto.Market
	partyBuf      PartyBuf
	partyByMarket map[string]map[string]struct{}
	mu            sync.Mutex
}

// NewParty instantiate a new party
func NewParty(log *logging.Logger, col Collateral, markets []proto.Market, partyBuf PartyBuf) *Party {
	partyByMarket := map[string]map[string]struct{}{}
	for _, v := range markets {
		partyByMarket[v.Id] = map[string]struct{}{}
	}
	return &Party{
		log:           log,
		collateral:    col,
		markets:       markets,
		partyBuf:      partyBuf,
		partyByMarket: partyByMarket,
	}
}

// GetForMarket returns the list of all the parties in a given market
func (p *Party) GetForMarket(mktID string) []string {
	parties := p.partyByMarket[mktID]
	out := make([]string, 0, len(parties))
	for k := range parties {
		out = append(out, k)
	}
	return out
}

func (p *Party) GetByMarketAndID(marketID, partyID string) (*types.Party, error) {
	if _, ok := p.partyByMarket[marketID][partyID]; ok {
		return &types.Party{Id: partyID}, nil
	}
	return nil, ErrPartyDoesNotExist
}

func (p *Party) addMarket(market proto.Market) {
	if _, found := p.partyByMarket[market.Id]; !found {
		p.mu.Lock()
		p.markets = append(p.markets, market)
		p.partyByMarket[market.Id] = map[string]struct{}{}
		p.mu.Unlock()
	}
}

func (p *Party) addParty(ptyID, mktID string) {
	p.partyByMarket[mktID][ptyID] = struct{}{}
}

// NotifyTraderAccountWithTopUpAmount will create a new party in the system
// and top-up it general account with the given amount
func (p *Party) NotifyTraderAccountWithTopUpAmount(
	notif *proto.NotifyTraderAccount, amount int64) error {
	return p.notifyTraderAccount(notif, amount)
}

// NotifyTraderAccount will create a new party in the system
// and top-up it general account with the default amount
func (p *Party) NotifyTraderAccount(notif *proto.NotifyTraderAccount) error {
	if notif.Amount == 0 {
		return p.notifyTraderAccount(notif, 1000000000000)
	}
	return p.notifyTraderAccount(notif, int64(notif.Amount))
}

func (p *Party) notifyTraderAccount(notif *proto.NotifyTraderAccount, amount int64) error {
	alreadyTopUp := map[string]struct{}{}

	// ignore errors as they can only happen when the party already exists
	p.partyBuf.Add(types.Party{Id: notif.TraderID})

	for _, mkt := range p.markets {
		p.addParty(notif.TraderID, mkt.Id)
		asset, err := mkt.GetAsset()
		if err != nil {
			p.log.Error("unable to get market asset",
				logging.Error(err))
			return err
		}

		// create account
		generalID := p.collateral.CreatePartyGeneralAccount(notif.TraderID, asset)
		if _, ok := alreadyTopUp[generalID]; !ok {
			alreadyTopUp[generalID] = struct{}{}
			// then credit the general account
			err = p.collateral.IncrementBalance(generalID, amount)
			if err != nil {
				p.log.Error("unable to topup trader account",
					logging.Error(err))
				return err
			}
			var acc *types.Account
			acc, err = p.collateral.GetAccountByID(generalID)
			if err != nil {
				p.log.Error("unable to get trader account",
					logging.String("party-id", notif.TraderID),
					logging.String("asset", asset),
					logging.Error(err))
				return err
			}
			if p.log.GetLevel() == logging.DebugLevel {
				p.log.Debug("party account top-up",
					logging.String("asset", asset),
					logging.String("party-id", notif.TraderID),
					logging.Int64("top-up-amount", amount),
					logging.Int64("new-balance", acc.Balance))
			}
		}
	}

	return nil
}
