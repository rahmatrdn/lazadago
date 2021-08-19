package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetFlexiComboDetailsSampleSkusResponseEntity struct{
    ProductId	int	`json:"product_id"`
    SkuId	int	`json:"sku_id"`
}
func (g GetFlexiComboDetailsSampleSkusResponseEntity) String() string {
    return lib.ObjectToString(g)
}
