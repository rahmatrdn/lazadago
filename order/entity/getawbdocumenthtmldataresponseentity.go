package entity

import (
    "github.com/alfarady/letgo/lib"
)

type GetAwbDocumentHtmlDataResponseEntity struct{
    Document	GetAwbDocumentHtmlDocumentResponseEntity	`json:"document"`
}
func (g GetAwbDocumentHtmlDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
