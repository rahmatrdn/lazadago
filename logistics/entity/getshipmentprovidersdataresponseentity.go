package entity

import (
    "github.com/alfarady/letgo/lib"
)

type GetShipmentProvidersDataResponseEntity struct{
    ShipmentProviders	[]GetShipmentProvidersShipmentProvidersResponseEntity	`json:"shipment_providers"`
}
func (g GetShipmentProvidersDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
