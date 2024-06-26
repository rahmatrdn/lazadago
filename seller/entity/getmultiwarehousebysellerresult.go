package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetMultiWarehouseBySellerResult struct {
	Result    GetMultiWarehouseBySellerResultResponseEntity `json:"result"`
	Type      string                                        `json:"type"`
	RequestId string                                        `json:"request_id"`
	Code      string                                        `json:"code"`
}

func (g GetMultiWarehouseBySellerResult) String() string {
	return lib.ObjectToString(g)
}

type GetMultiWarehouseBySellerDetailResponseEntity struct{}
