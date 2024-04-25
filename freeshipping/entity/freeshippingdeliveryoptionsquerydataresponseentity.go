package entity

import (
    "github.com/alfarady/letgo/lib"
)

type FreeShippingDeliveryOptionsQueryDataResponseEntity struct{
    Name	string	`json:"name"`
    Value	string	`json:"value"`
}
func (g FreeShippingDeliveryOptionsQueryDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
