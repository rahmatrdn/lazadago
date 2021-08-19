package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetMultipleOrderItemsDataResponseEntity struct{
    OrderItems	[]GetMultipleOrderItemsOrderItemsResponseEntity	`json:"order_items"`
    OrderNumber	int	`json:"order_number"`
    OrderId	int	`json:"order_id"`
}
func (g GetMultipleOrderItemsDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
