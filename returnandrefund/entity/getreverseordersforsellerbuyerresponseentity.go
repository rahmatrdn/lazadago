package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetReverseOrdersForSellerBuyerResponseEntity struct{
    BuyerId	int	`json:"buyer_id"`
}
func (g GetReverseOrdersForSellerBuyerResponseEntity) String() string {
    return lib.ObjectToString(g)
}
