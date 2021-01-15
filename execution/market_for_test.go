package execution

import types "code.vegaprotocol.io/vega/proto"

// GetPeggedOrderCount returns the number of pegged orders in the market
func (m *Market) GetPeggedOrderCount() int {
	return len(m.peggedOrders)
}

// GetParkedOrderCount returns hte number of parked orders in the market
func (m *Market) GetParkedOrderCount() int {
	var count int
	for _, order := range m.peggedOrders {
		if order.Status == types.Order_STATUS_PARKED {
			count++
		}
	}
	return count
}

// GetPeggedExpiryOrderCount returns the number of pegged order that can expire
func (m *Market) GetPeggedExpiryOrderCount() int {
	return m.expiringOrders.GetExpiryingOrderCount()
}

// TSCalc returns the local tsCalc instance
func (m *Market) TSCalc() TargetStakeCalculator {
	return m.tsCalc
}

func (m *Market) State() types.Market_State {
	return m.mkt.State
}
