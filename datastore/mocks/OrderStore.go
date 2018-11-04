// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import filters "vega/filters"
import mock "github.com/stretchr/testify/mock"
import msg "vega/msg"

// OrderStore is an autogenerated mock type for the OrderStore type
type OrderStore struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *OrderStore) Close() {
	_m.Called()
}

// Delete provides a mock function with given fields: order
func (_m *OrderStore) Delete(order *msg.Order) error {
	ret := _m.Called(order)

	var r0 error
	if rf, ok := ret.Get(0).(func(*msg.Order) error); ok {
		r0 = rf(order)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByMarket provides a mock function with given fields: market, _a1
func (_m *OrderStore) GetByMarket(market string, _a1 *filters.OrderQueryFilters) ([]*msg.Order, error) {
	ret := _m.Called(market, _a1)

	var r0 []*msg.Order
	if rf, ok := ret.Get(0).(func(string, *filters.OrderQueryFilters) []*msg.Order); ok {
		r0 = rf(market, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*msg.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *filters.OrderQueryFilters) error); ok {
		r1 = rf(market, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByMarketAndId provides a mock function with given fields: market, id
func (_m *OrderStore) GetByMarketAndId(market string, id string) (*msg.Order, error) {
	ret := _m.Called(market, id)

	var r0 *msg.Order
	if rf, ok := ret.Get(0).(func(string, string) *msg.Order); ok {
		r0 = rf(market, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*msg.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(market, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByParty provides a mock function with given fields: party, _a1
func (_m *OrderStore) GetByParty(party string, _a1 *filters.OrderQueryFilters) ([]*msg.Order, error) {
	ret := _m.Called(party, _a1)

	var r0 []*msg.Order
	if rf, ok := ret.Get(0).(func(string, *filters.OrderQueryFilters) []*msg.Order); ok {
		r0 = rf(party, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*msg.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *filters.OrderQueryFilters) error); ok {
		r1 = rf(party, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByPartyAndId provides a mock function with given fields: party, id
func (_m *OrderStore) GetByPartyAndId(party string, id string) (*msg.Order, error) {
	ret := _m.Called(party, id)

	var r0 *msg.Order
	if rf, ok := ret.Get(0).(func(string, string) *msg.Order); ok {
		r0 = rf(party, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*msg.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(party, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMarketDepth provides a mock function with given fields: market
func (_m *OrderStore) GetMarketDepth(market string) (*msg.MarketDepth, error) {
	ret := _m.Called(market)

	var r0 *msg.MarketDepth
	if rf, ok := ret.Get(0).(func(string) *msg.MarketDepth); ok {
		r0 = rf(market)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*msg.MarketDepth)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(market)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Notify provides a mock function with given fields:
func (_m *OrderStore) Notify() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Post provides a mock function with given fields: order
func (_m *OrderStore) Post(order *msg.Order) error {
	ret := _m.Called(order)

	var r0 error
	if rf, ok := ret.Get(0).(func(*msg.Order) error); ok {
		r0 = rf(order)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PostBatch provides a mock function with given fields: batch
func (_m *OrderStore) PostBatch(batch []*msg.Order) error {
	ret := _m.Called(batch)

	var r0 error
	if rf, ok := ret.Get(0).(func([]*msg.Order) error); ok {
		r0 = rf(batch)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Put provides a mock function with given fields: order
func (_m *OrderStore) Put(order *msg.Order) error {
	ret := _m.Called(order)

	var r0 error
	if rf, ok := ret.Get(0).(func(*msg.Order) error); ok {
		r0 = rf(order)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Subscribe provides a mock function with given fields: orders
func (_m *OrderStore) Subscribe(orders chan<- []msg.Order) uint64 {
	ret := _m.Called(orders)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(chan<- []msg.Order) uint64); ok {
		r0 = rf(orders)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// Unsubscribe provides a mock function with given fields: id
func (_m *OrderStore) Unsubscribe(id uint64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
