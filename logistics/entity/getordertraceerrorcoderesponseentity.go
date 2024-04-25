package entity

import (
    "github.com/alfarady/letgo/lib"
)

type GetOrderTraceErrorCodeResponseEntity struct{
    DisplayMessage	string	`json:"displayMessage"`
}
func (g GetOrderTraceErrorCodeResponseEntity) String() string {
    return lib.ObjectToString(g)
}
