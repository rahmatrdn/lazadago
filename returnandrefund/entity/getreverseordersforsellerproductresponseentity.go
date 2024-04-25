package entity

import (
    "github.com/alfarady/letgo/lib"
)

type GetReverseOrdersForSellerProductResponseEntity struct{
    ProductId	int	`json:"product_id"`
    ProductSku	string	`json:"product_sku"`
}
func (g GetReverseOrdersForSellerProductResponseEntity) String() string {
    return lib.ObjectToString(g)
}
