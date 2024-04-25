package entity

import (
    "github.com/alfarady/letgo/lib"
)

type SetStatusToSOFDeliveredOrderItemsResponseEntity struct{
    OrderItemId	int64	`json:"order_item_id"`
    PurchaseOrderId	int64	`json:"purchase_order_id"`
    PurchaseOrderNumber	string	`json:"purchase_order_number"`
}
func (g SetStatusToSOFDeliveredOrderItemsResponseEntity) String() string {
    return lib.ObjectToString(g)
}
