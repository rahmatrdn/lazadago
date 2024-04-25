package entity

import (
    "github.com/alfarady/letgo/lib"
)

type DataMoatLoginResultResponseEntity struct{
    Msg	string	`json:"msg"`
    Success	bool	`json:"success"`
}
func (g DataMoatLoginResultResponseEntity) String() string {
    return lib.ObjectToString(g)
}
