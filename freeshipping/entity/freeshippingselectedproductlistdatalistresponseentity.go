package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type FreeShippingSelectedProductListDataListResponseEntity struct{
    ProductId	int	`json:"product_id"`
    SkuIds	[]int	`json:"sku_ids"`
}
func (g FreeShippingSelectedProductListDataListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
