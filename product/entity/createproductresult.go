package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type CreateProductResult struct {
	Data      CreateProductDataResponseEntity     `json:"data"`
	Type      string                              `json:"type"`
	Code      string                              `json:"code"`
	Message   string                              `json:"message"`
	RequestId string                              `json:"request_id"`
	Detail    []CreateProductDetailResponseEntity `json:"detail"`
}

func (g CreateProductResult) String() string {
	return lib.ObjectToString(g)
}

type CreateProductDetailResponseEntity struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Code    string `json:"code"`
}
