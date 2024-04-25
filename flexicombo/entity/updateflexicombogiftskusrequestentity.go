package entity

import (
    "github.com/alfarady/letgo/lib"
)

type UpdateFlexiComboGiftSkusRequestEntity struct{
    ProductId	int64	`json:"productId"`
    SkuId	int64	`json:"skuId"`
}
func (g UpdateFlexiComboGiftSkusRequestEntity) String() string {
    return lib.ObjectToString(g)
}
