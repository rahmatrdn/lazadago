package entity

import (
    "github.com/alfarady/letgo/lib"
)

type CreateFlexiComboSampleSkusRequestEntity struct{
    ProductId	int64	`json:"productId"`
    SkuId	int64	`json:"skuId"`
}
func (g CreateFlexiComboSampleSkusRequestEntity) String() string {
    return lib.ObjectToString(g)
}
