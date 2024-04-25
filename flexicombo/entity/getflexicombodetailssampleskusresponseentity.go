package entity

import (
    "github.com/alfarady/letgo/lib"
)

type GetFlexiComboDetailsSampleSkusResponseEntity struct{
    ProductId	int64	`json:"product_id"`
    SkuId	int64	`json:"sku_id"`
}
func (g GetFlexiComboDetailsSampleSkusResponseEntity) String() string {
    return lib.ObjectToString(g)
}
