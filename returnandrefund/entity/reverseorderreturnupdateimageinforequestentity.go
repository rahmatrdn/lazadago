package entity

import (
    "github.com/alfarady/letgo/lib"
)

type ReverseOrderReturnUpdateImageInfoRequestEntity struct{
    Name	string	`json:"name"`
    Url	string	`json:"url"`
}
func (g ReverseOrderReturnUpdateImageInfoRequestEntity) String() string {
    return lib.ObjectToString(g)
}
