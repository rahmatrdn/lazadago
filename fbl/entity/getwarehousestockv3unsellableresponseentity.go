package entity

import (
    "github.com/alfarady/letgo/lib"
)

type GetWarehouseStockV3UnsellableResponseEntity struct{
    Available	int	`json:"available"`
    Reserved	int	`json:"reserved"`
}
func (g GetWarehouseStockV3UnsellableResponseEntity) String() string {
    return lib.ObjectToString(g)
}
