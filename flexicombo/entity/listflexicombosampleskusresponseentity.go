package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type ListFlexiComboSampleSkusResponseEntity struct{
    ProductId	int	`json:"product_id"`
    SkuId	int	`json:"sku_id"`
}
func (g ListFlexiComboSampleSkusResponseEntity) String() string {
    return lib.ObjectToString(g)
}
