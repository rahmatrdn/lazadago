package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type ListFlexiComboGiftSkusResponseEntity struct{
    ProductId	int	`json:"product_id"`
    SkuId	int	`json:"sku_id"`
}
func (g ListFlexiComboGiftSkusResponseEntity) String() string {
    return lib.ObjectToString(g)
}
