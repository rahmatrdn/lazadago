package entity

import (
    "github.com/alfarady/letgo/lib"
)

type LazadaBigbagCancelUserInfoRequestEntity struct{
    AppUserKey	string	`json:"appUserKey"`
}
func (g LazadaBigbagCancelUserInfoRequestEntity) String() string {
    return lib.ObjectToString(g)
}
