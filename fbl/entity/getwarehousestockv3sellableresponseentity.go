package entity

import (
    "github.com/alfarady/letgo/lib"
)

type GetWarehouseStockV3SellableResponseEntity struct{
    Available	int	`json:"available"`
    Reserved	int	`json:"reserved"`
}
func (g GetWarehouseStockV3SellableResponseEntity) String() string {
    return lib.ObjectToString(g)
}
