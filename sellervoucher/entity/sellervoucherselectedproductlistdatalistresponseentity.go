package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type SellerVoucherSelectedProductListDataListResponseEntity struct{
    ProductId	int	`json:"product_id"`
    SkuIds	[]int	`json:"sku_ids"`
}
func (g SellerVoucherSelectedProductListDataListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
