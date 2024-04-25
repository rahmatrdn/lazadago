package entity

import (
    "github.com/alfarady/letgo/lib"
)

type FreeShippingListTiersResponseEntity struct{
    Filter	string	`json:"filter"`
    Result	string	`json:"result"`
}
func (g FreeShippingListTiersResponseEntity) String() string {
    return lib.ObjectToString(g)
}
