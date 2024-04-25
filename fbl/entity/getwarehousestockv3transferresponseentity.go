package entity

import (
    "github.com/alfarady/letgo/lib"
)

type GetWarehouseStockV3TransferResponseEntity struct{
    Available	int	`json:"available"`
    Reserved	int	`json:"reserved"`
}
func (g GetWarehouseStockV3TransferResponseEntity) String() string {
    return lib.ObjectToString(g)
}
