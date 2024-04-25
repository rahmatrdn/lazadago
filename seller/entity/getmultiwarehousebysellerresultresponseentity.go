package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetMultiWarehouseBySellerResultResponseEntity struct {
	Success    bool                                             `json:"success"`
	NotSuccess bool                                             `json:"not_success"`
	Module     GetMultiWarehouseBySellerModuleResponseEntity    `json:"module"`
	ErrorCode  GetMultiWarehouseBySellerErrorCodeResponseEntity `json:"error_code"`
}

func (g GetMultiWarehouseBySellerResultResponseEntity) String() string {
	return lib.ObjectToString(g)
}

type GetMultiWarehouseBySellerModuleResponseEntity struct {
	Country       string `json:"country"`
	DefaultAdress string `json:"default_address"`
	Province      string `json:"province"`
	City          string `json:"city"`
	DetailAddress string `json:"detail_address"`
	WHCode        string `json:"warehouse_code"`
	District      string `json:"district"`
	PostCode      string `json:"post_code"`
	Name          string `json:"name"`
	Status        string `json:"status"`
}
