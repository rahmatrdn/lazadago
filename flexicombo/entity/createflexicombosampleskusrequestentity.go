package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type CreateFlexiComboSampleSkusRequestEntity struct{
    ProductId	int	`json:"productId"`
    SkuId	int	`json:"skuId"`
}
func (g CreateFlexiComboSampleSkusRequestEntity) String() string {
    return lib.ObjectToString(g)
}
