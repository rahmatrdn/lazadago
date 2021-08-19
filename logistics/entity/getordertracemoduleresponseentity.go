package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetOrderTraceModuleResponseEntity struct{
    WarehouseDetailInfo	string	`json:"warehouse_detail_info"`
    OfcOrderId	string	`json:"ofc_order_id"`
    PackageDetailInfoList	[]GetOrderTracePackageDetailInfoListResponseEntity	`json:"package_detail_info_list"`
}
func (g GetOrderTraceModuleResponseEntity) String() string {
    return lib.ObjectToString(g)
}
