package entity

import (
    "github.com/alfarady/letgo/lib"
)

type GetOVOOrdersResultResponseEntity struct{
    ErrorCode	string	`json:"errorCode"`
    Success	string	`json:"success"`
    TradeOrders	[]GetOVOOrdersTradeOrdersResponseEntity	`json:"tradeOrders"`
}
func (g GetOVOOrdersResultResponseEntity) String() string {
    return lib.ObjectToString(g)
}
