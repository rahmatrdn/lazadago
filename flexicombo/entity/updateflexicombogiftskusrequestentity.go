package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UpdateFlexiComboGiftSkusRequestEntity struct{
    ProductId	int	`json:"productId"`
    SkuId	int	`json:"skuId"`
}
func (g UpdateFlexiComboGiftSkusRequestEntity) String() string {
    return lib.ObjectToString(g)
}
