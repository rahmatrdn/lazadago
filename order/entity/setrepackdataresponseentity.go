package entity

import (
    "github.com/alfarady/letgo/lib"
)

type SetRepackDataResponseEntity struct{
    PackageId	string	`json:"package_id"`
}
func (g SetRepackDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
