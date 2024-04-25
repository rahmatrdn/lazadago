package entity

import (
    "github.com/alfarady/letgo/lib"
)

type GetMultipleOrderItemsDataResponseEntity struct{
    OrderItems	[]GetMultipleOrderItemsOrderItemsResponseEntity	`json:"order_items"`
    OrderNumber	int64	`json:"order_number"`
    OrderId	int64	`json:"order_id"`
}
func (g GetMultipleOrderItemsDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
