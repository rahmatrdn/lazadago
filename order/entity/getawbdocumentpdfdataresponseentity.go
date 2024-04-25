package entity

import (
    "github.com/alfarady/letgo/lib"
)

type GetAwbDocumentPDFDataResponseEntity struct{
    Document	GetAwbDocumentPDFDocumentResponseEntity	`json:"document"`
}
func (g GetAwbDocumentPDFDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
