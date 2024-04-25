package entity

import (
    "github.com/alfarady/letgo/lib"
)

type ListFlexiComboProductsDataResponseEntity struct{
    Total	int	`json:"total"`
    Current	int	`json:"current"`
    DataList	[]int	`json:"data_list"`
    PageSize	int	`json:"page_size"`
}
func (g ListFlexiComboProductsDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
