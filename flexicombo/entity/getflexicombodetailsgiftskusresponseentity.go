package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetFlexiComboDetailsGiftSkusResponseEntity struct{
    ProductId	int	`json:"product_id"`
    SkuId	int	`json:"sku_id"`
}
func (g GetFlexiComboDetailsGiftSkusResponseEntity) String() string {
    return lib.ObjectToString(g)
}
