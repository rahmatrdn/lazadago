package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetReverseOrderDetailDataResponseEntity struct{
    ReverseOrderId	int	`json:"reverse_order_id"`
    TradeOrderId	int	`json:"trade_order_id"`
    RequestType	string	`json:"request_type"`
    ShippingType	string	`json:"shipping_type"`
    IsRtm	bool	`json:"is_rtm"`
    ReverseOrderLineDTOList	[]GetReverseOrderDetailReverseOrderLineDTOListResponseEntity	`json:"reverseOrderLineDTOList"`
}
func (g GetReverseOrderDetailDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
