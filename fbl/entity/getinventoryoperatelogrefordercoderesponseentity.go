package entity

import (
    "github.com/alfarady/letgo/lib"
)

type GetInventoryOperateLogRefOrderCodeResponseEntity struct{
    Type	string	`json:"type"`
    OrderCode	string	`json:"order_code"`
}
func (g GetInventoryOperateLogRefOrderCodeResponseEntity) String() string {
    return lib.ObjectToString(g)
}
