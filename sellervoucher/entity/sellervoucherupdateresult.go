package entity

import (
    "github.com/wjp-letgo/letgo/lib"
)

type SellerVoucherUpdateResult struct{
    Data	int	`json:"data"`
    Success	bool	`json:"success"`
    ErrorCode	int	`json:"error_code"`
    ErrorMsg	string	`json:"error_msg"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]SellerVoucherUpdateDetailResponseEntity	`json:"detail"`
}
func (g SellerVoucherUpdateResult) String() string {
    return lib.ObjectToString(g)
}
type SellerVoucherUpdateDetailResponseEntity struct{}
