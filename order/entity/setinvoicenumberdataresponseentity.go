package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type SetInvoiceNumberDataResponseEntity struct{
    OrderItemId	int	`json:"order_item_id"`
    InvoiceNumber	string	`json:"invoice_number"`
}
func (g SetInvoiceNumberDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
