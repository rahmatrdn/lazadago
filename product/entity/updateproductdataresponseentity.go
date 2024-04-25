package entity

import (
    "github.com/alfarady/letgo/lib"
)

type UpdateProductDataResponseEntity struct{
    Variation	UpdateProductVariationResponseEntity	`json:"variation"`
}
func (g UpdateProductDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
