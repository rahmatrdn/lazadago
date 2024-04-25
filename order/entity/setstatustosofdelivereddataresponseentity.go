package entity

import (
    "github.com/alfarady/letgo/lib"
)

type SetStatusToSOFDeliveredDataResponseEntity struct{
    OrderItems	[]SetStatusToSOFDeliveredOrderItemsResponseEntity	`json:"order_items"`
}
func (g SetStatusToSOFDeliveredDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
