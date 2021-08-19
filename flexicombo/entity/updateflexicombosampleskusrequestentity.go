package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UpdateFlexiComboSampleSkusRequestEntity struct{
    ProductId	int	`json:"productId"`
    SkuId	int	`json:"skuId"`
}
func (g UpdateFlexiComboSampleSkusRequestEntity) String() string {
    return lib.ObjectToString(g)
}
