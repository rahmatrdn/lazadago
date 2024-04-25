package entity

import (
    "github.com/alfarady/letgo/lib"
)

type GetReverseOrderDetailProductDTOResponseEntity struct{
    ProductId	int	`json:"product_id"`
    Sku	string	`json:"sku"`
}
func (g GetReverseOrderDetailProductDTOResponseEntity) String() string {
    return lib.ObjectToString(g)
}
