package entity

import (
    "github.com/alfarady/letgo/lib"
)

type GlobalEticketMerchantMaSendIsvMaListRequestEntity struct{
    Code	string	`json:"code"`
    Num	int	`json:"num"`
}
func (g GlobalEticketMerchantMaSendIsvMaListRequestEntity) String() string {
    return lib.ObjectToString(g)
}
