package entity

import (
    "github.com/alfarady/letgo/lib"
)

type QueryAddressInformaitonUserInfoRequestEntity struct{
    AppUserKey	string	`json:"appUserKey"`
}
func (g QueryAddressInformaitonUserInfoRequestEntity) String() string {
    return lib.ObjectToString(g)
}
