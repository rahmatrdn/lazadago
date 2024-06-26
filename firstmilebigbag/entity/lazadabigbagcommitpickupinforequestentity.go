package entity

import (
    "github.com/wjp-letgo/letgo/lib"
)

type LazadaBigbagCommitPickupInfoRequestEntity struct{
    Address	LazadaBigbagCommitAddressRequestEntity	`json:"address"`
    Phone	string	`json:"phone"`
    Name	string	`json:"name"`
    Mobile	string	`json:"mobile"`
    Email	string	`json:"email"`
    AddressId	int	`json:"addressId"`
}
func (g LazadaBigbagCommitPickupInfoRequestEntity) String() string {
    return lib.ObjectToString(g)
}
