package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type CreateFlexiComboGiftSkusRequestEntity struct{
    ProductId	int	`json:"productId"`
    SkuId	int	`json:"skuId"`
}
func (g CreateFlexiComboGiftSkusRequestEntity) String() string {
    return lib.ObjectToString(g)
}
