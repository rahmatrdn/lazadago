package entity

import (
    "github.com/alfarady/letgo/lib"
)

type FreeShippingGetTiersResponseEntity struct{
    Filter	string	`json:"filter"`
    Result	string	`json:"result"`
}
func (g FreeShippingGetTiersResponseEntity) String() string {
    return lib.ObjectToString(g)
}
