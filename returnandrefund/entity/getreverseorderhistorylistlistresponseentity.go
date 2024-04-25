package entity

import (
    "github.com/alfarady/letgo/lib"
)

type GetReverseOrderHistoryListListResponseEntity struct{
    Operator	string	`json:"operator"`
    Picture	[]string	`json:"picture"`
    Time	int	`json:"time"`
}
func (g GetReverseOrderHistoryListListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
