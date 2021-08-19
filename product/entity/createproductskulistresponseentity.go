package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type CreateProductSkuListResponseEntity struct{
    SellerSku	string	`json:"seller_sku"`
    ShopSku	string	`json:"shop_sku"`
    SkuId	int	`json:"sku_id"`
}
func (g CreateProductSkuListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
