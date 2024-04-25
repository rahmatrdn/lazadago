package entity

import (
    "github.com/alfarady/letgo/lib"
)

type ReverseOrderCancelValidateReasonOptionsResponseEntity struct{
    ReasonName	string	`json:"reason_name"`
    ReasonId	string	`json:"reason_id"`
}
func (g ReverseOrderCancelValidateReasonOptionsResponseEntity) String() string {
    return lib.ObjectToString(g)
}
