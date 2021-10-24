package orderbook

type Orderbook struct {
	restingOrders []*Order
}

func New() *Orderbook {
	return &Orderbook{
        restingOrders: make([]*Order, 0),
    }
}

func (o *Orderbook) otherSideROrders(order *Order) []*Order {
    otherSROrders := make([]*Order, 0)
    for _, restingOrder := range o.restingOrders {
        if order.Side == restingOrder.Side {
            continue
        }
        otherSROrders = append(otherSROrders, restingOrder)
    }
    return otherSROrders
}

func (o *Orderbook) deleteOrder(restbookOrder *Order) {
    for i, ord := range o.restingOrders {
        if ord == restbookOrder {
            o.restingOrders = append(o.restingOrders[:i], o.restingOrders[i+1:]...)
            return
        }
    }
}

func (o *Orderbook) executeOrders(order, restingOrder *Order) (trade *Trade, fulfilled bool) {
    trade = &Trade{}

    switch order.Side {
    case SideAsk:
        trade.Ask = order
        trade.Bid = restingOrder
    case SideBid:
        trade.Ask = restingOrder
        trade.Bid = order
    }

    switch {
    case order.Volume < restingOrder.Volume:
        trade.Volume = order.Volume
        trade.Price = restingOrder.Price
        fulfilled = true
        restingOrder.Volume = restingOrder.Volume - order.Volume
        order.Volume = 0
    case order.Volume > restingOrder.Volume:
        trade.Volume = restingOrder.Volume
        trade.Price = restingOrder.Price
        fulfilled = false
        order.Volume = order.Volume - restingOrder.Volume
        restingOrder.Volume = 0
        o.deleteOrder(restingOrder)
    default:
        trade.Volume = order.Volume
        trade.Price = restingOrder.Price
        fulfilled = true
        order.Volume = 0
        restingOrder.Volume = 0
        o.deleteOrder(restingOrder)
    }
    return
}

func (o *Orderbook) addToRestingOrders(order *Order) {
    o.restingOrders = append(o.restingOrders, order)
}

func (o *Orderbook) findMostProfitableLimit(order *Order) *Order {
    possibleOrders := o.otherSideROrders(order)
    if len(possibleOrders) == 0 {
        return nil
    }
    var mostProfitable *Order
    for _, restingOrder := range possibleOrders {
        switch order.Side {
        case SideAsk:
            if order.Price <= restingOrder.Price &&
            (mostProfitable == nil ||
            restingOrder.Price - order.Price > mostProfitable.Price - order.Price) {
                mostProfitable = restingOrder
            }
        case SideBid:
            if order.Price >= restingOrder.Price &&
            (mostProfitable == nil ||
            order.Price - restingOrder.Price > order.Price - mostProfitable.Price) {
                mostProfitable = restingOrder
            }
        }
    }
    return mostProfitable
}

func (o *Orderbook) ProcessLimitOrder(order *Order) (trades []*Trade) {
    trades = make([]*Trade, 0)

    for profitable := o.findMostProfitableLimit(order); profitable != nil; profitable = o.findMostProfitableLimit(order) {
        trade, fulfilled := o.executeOrders(order, profitable)
        trades = append(trades, trade)
        if fulfilled {
            return
        }
    }
    return
}

func (o *Orderbook) findMostProfitableMarket(order *Order) *Order {
    possibleOrders := o.otherSideROrders(order)
    if len(possibleOrders) == 0 {
        return nil
    }
    var mostProfitable *Order
    for _, restingOrder := range possibleOrders {
        switch order.Side {
        case SideAsk:
            if mostProfitable == nil ||
            restingOrder.Price > mostProfitable.Price {
                mostProfitable = restingOrder
            }
        case SideBid:
            if mostProfitable == nil ||
            restingOrder.Price < mostProfitable.Price {
                mostProfitable = restingOrder
            }
        }
    }
    return mostProfitable
}

func (o *Orderbook) ProcessMarketOrder(order *Order) (trades []*Trade) {
    trades = make([]*Trade, 0)

    for profitable := o.findMostProfitableMarket(order); profitable != nil; profitable = o.findMostProfitableMarket(order) {
        trade, fulfilled := o.executeOrders(order, profitable)
        trades = append(trades, trade)
        if fulfilled {
            return
        }
    }
    return
}

func (o *Orderbook) Match(order *Order) (trades []*Trade, rejected *Order) {
    // try to fulfill the order
    // if not -- mark it as resting order
    // else -- alter other orders due to fulfillment of object order
    // in trades our alter info resides, second arg is reject order
    // it may reject order only and only if order is market order and there are no
    // resting orders in orderbook that can fulfill it
   
    switch order.Kind {
    case KindLimit:
        trades = o.ProcessLimitOrder(order)
        if order.Volume > 0 {
            o.addToRestingOrders(order)
        }
    case KindMarket:
        trades = o.ProcessMarketOrder(order)
        if order.Volume > 0 {
            rejected = order
        }
    }

	return
}
