package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetUnfilledAttributeProductsResponseEntity struct{
    ItemId	int	`json:"item_id"`
    PrimaryCategory	int	`json:"primary_category"`
    SellerSku	string	`json:"seller_sku"`
    Attributes	[]GetUnfilledAttributeAttributesResponseEntity	`json:"attributes"`
}
func (g GetUnfilledAttributeProductsResponseEntity) String() string {
    return lib.ObjectToString(g)
}
