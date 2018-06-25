package datastore

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewMemStore_ReturnsNewMemStoreInstance(t *testing.T) {
	var memStore = NewMemStore()
	assert.NotNil(t, memStore)
}

func TestNewTradeStore_ReturnsNewTradeStoreInstance(t *testing.T) {
	var memStore = NewMemStore()
	var newTradeStore = NewTradeStore(&memStore)
	assert.NotNil(t, newTradeStore)
}

func TestNewOrderStore_ReturnsNewOrderStoreInstance(t *testing.T) {
	var memStore = NewMemStore()
	var newOrderStore = NewOrderStore(&memStore)
	assert.NotNil(t, newOrderStore)
}

func TestMemOrderStore_PutAndGetNewOrder(t *testing.T) {
	var memStore = NewMemStore()
	var newOrderStore = NewOrderStore(&memStore)

	var order = Order{
		ID: "45305210ff7a9bb9450b1833cc10368a",
	}

	err := newOrderStore.Put(&order)
	assert.Nil(t, err)

	o, err := newOrderStore.Get(order.ID)
	assert.Nil(t, err)
	assert.Equal(t, &order, o)
}


func TestMemOrderStore_PutAndGetExistingOrder(t *testing.T) {
	var memStore = NewMemStore()
	var newOrderStore = NewOrderStore(&memStore)

	var order = Order{
		ID:        "c471bdd5f381aa3654d98f4591eaa968",
		Market:    "market",
		Party:     "tester",
		Price:     100,
		Size:      1,
	}

	err := newOrderStore.Put(&order)
	assert.Nil(t, err)

	o, err := newOrderStore.Get(order.ID)
	assert.Nil(t, err)
	assert.Equal(t, uint64(100), o.Price)
	assert.Equal(t, uint64(1), o.Size)

	order.Price = 1000
	order.Size = 5

	err = newOrderStore.Put(&order)
	assert.Nil(t, err)

	o, err = newOrderStore.Get(order.ID)
	assert.Nil(t, err)
	assert.Equal(t, &order, o)
	assert.Equal(t, uint64(1000), o.Price)
	assert.Equal(t, uint64(5), o.Size)
}



func TestMemOrderStore_PutAndDeleteOrder(t *testing.T) {
	var memStore = NewMemStore()
	var newOrderStore = NewOrderStore(&memStore)

	var order = Order{
		ID: "45305210ff7a9bb9450b1833cc10368a",
	}

	err := newOrderStore.Put(&order)
	assert.Nil(t, err)

	o, err := newOrderStore.Get(order.ID)
	assert.Nil(t, err)
	assert.Equal(t, &order, o)

	err = newOrderStore.Delete(o)
	assert.Nil(t, err)

	o, err = newOrderStore.Get(order.ID)
	assert.Nil(t, o)
}

func TestMemOrderStore_PutAndGetTrade(t *testing.T) {
	tests := []struct {
		name string
		in Trade
		out Trade
	}{
		{"Empty trade", Trade { OrderID: "d41d8cd98f00b204e9800998ecf8427e" }, Trade { OrderID: "d41d8cd98f00b204e9800998ecf8427e" }, },
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var memStore = NewMemStore()
			var newOrderStore = NewOrderStore(&memStore)
			var newTradeStore = NewTradeStore(&memStore)

			var order = Order{
				ID: "d41d8cd98f00b204e9800998ecf8427e",
			}

			err := newOrderStore.Put(&order)
			assert.Nil(t, err)

			err = newTradeStore.Put(&test.in)
			assert.Nil(t, err)

			tr, err := newTradeStore.Get(test.in.ID)
			assert.Equal(t, &test.out, tr)
		})
	}
}


func TestMemOrderStore_PutAndDeleteTrade(t *testing.T) {
	var memStore = NewMemStore()
	var newOrderStore = NewOrderStore(&memStore)
	var newTradeStore = NewTradeStore(&memStore)

	var order = Order{
		ID: "d41d8cd98f00b204e9800998ecf8427e",
	}
	var trade = Trade{
		OrderID: "d41d8cd98f00b204e9800998ecf8427e",
	}

	err := newOrderStore.Put(&order)
	assert.Nil(t, err)

	err = newTradeStore.Put(&trade)
	assert.Nil(t, err)

	tr, err := newTradeStore.Get(trade.ID)
	assert.Equal(t, &trade, tr)

	err = newTradeStore.Delete(tr)
	assert.Nil(t, err)

	tr, err = newTradeStore.Get(trade.ID)
	assert.Nil(t, tr)
}

func TestMemOrderStore_PutTradeOrderNotFound(t *testing.T) {
	var memStore = NewMemStore()
	var newTradeStore = NewTradeStore(&memStore)
	trade := Trade {
		ID: "one",
		OrderID: "mystery",
	}
	err := newTradeStore.Put(&trade)
	assert.Error(t, err)
}

func TestMemOrderStore_PutAndFindByOrderId(t *testing.T) {
	var memStore = NewMemStore()
	var newOrderStore = NewOrderStore(&memStore)
	var newTradeStore = NewTradeStore(&memStore)

	trade1 := Trade {
	   ID: "one",
	   OrderID: "d41d8cd98f00b204e9800998ecf8427e",
	}
	trade2 := Trade {
		ID: "two",
		OrderID: "d41d8cd98f00b204e9800998ecf8427e",
	}
	order := Order{
		ID: "d41d8cd98f00b204e9800998ecf8427e",
	}

	err := newOrderStore.Put(&order)
	assert.Nil(t, err)

	err = newTradeStore.Put(&trade1)
	assert.Nil(t, err)
	
	err = newTradeStore.Put(&trade2)
	assert.Nil(t, err)

	trades, err := newTradeStore.FindByOrderID(order.ID)
	assert.Nil(t, err)

	assert.Equal(t, 2, len(trades))
	assert.Equal(t, "one", trades[0].ID)
	assert.Equal(t, "two", trades[1].ID)
}