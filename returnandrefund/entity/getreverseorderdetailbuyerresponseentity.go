package entity

import (
    "github.com/alfarady/letgo/lib"
)

type GetReverseOrderDetailBuyerResponseEntity struct{
    UserId	int	`json:"user_id"`
}
func (g GetReverseOrderDetailBuyerResponseEntity) String() string {
    return lib.ObjectToString(g)
}
