package entity

import (
    "github.com/wjp-letgo/letgo/lib"
)

type QueryLazadaBigbagInfoUserInfoRequestEntity struct{
    AppUserKey	string	`json:"appUserKey"`
}
func (g QueryLazadaBigbagInfoUserInfoRequestEntity) String() string {
    return lib.ObjectToString(g)
}
