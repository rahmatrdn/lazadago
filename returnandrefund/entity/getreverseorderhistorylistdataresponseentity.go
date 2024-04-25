package entity

import (
    "github.com/alfarady/letgo/lib"
)

type GetReverseOrderHistoryListDataResponseEntity struct{
    List	[]GetReverseOrderHistoryListListResponseEntity	`json:"list"`
    PageInfo	GetReverseOrderHistoryListPageInfoResponseEntity	`json:"page_info"`
}
func (g GetReverseOrderHistoryListDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
