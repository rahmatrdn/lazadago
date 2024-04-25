package entity

import (
    "github.com/alfarady/letgo/lib"
)

type SetStatusToPackedByMarketplaceDataResponseEntity struct{
    OrderItems	[]SetStatusToPackedByMarketplaceOrderItemsResponseEntity	`json:"order_items"`
}
func (g SetStatusToPackedByMarketplaceDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
