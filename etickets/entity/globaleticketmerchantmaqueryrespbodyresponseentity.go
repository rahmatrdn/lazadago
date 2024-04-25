package entity

import (
    "github.com/alfarady/letgo/lib"
)

type GlobalEticketMerchantMaQueryRespBodyResponseEntity struct{
    Certificate	GlobalEticketMerchantMaQueryCertificateResponseEntity	`json:"certificate"`
}
func (g GlobalEticketMerchantMaQueryRespBodyResponseEntity) String() string {
    return lib.ObjectToString(g)
}
