package entity

import (
    "github.com/alfarady/letgo/lib"
)

type SetRepackResult struct{
    Data	SetRepackDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]SetRepackDetailResponseEntity	`json:"detail"`
}
func (g SetRepackResult) String() string {
    return lib.ObjectToString(g)
}

type SetRepackDetailResponseEntity struct{
    
}
