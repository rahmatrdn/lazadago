package entity

import (
    "github.com/alfarady/letgo/lib"
)

type GetLazadaBigbagPDFLableUserInfoRequestEntity struct{
    AppUserKey	string	`json:"appUserKey"`
}
func (g GetLazadaBigbagPDFLableUserInfoRequestEntity) String() string {
    return lib.ObjectToString(g)
}
