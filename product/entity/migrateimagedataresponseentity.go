package entity

import (
    "github.com/alfarady/letgo/lib"
)

type MigrateImageDataResponseEntity struct{
    Image	MigrateImageImageResponseEntity	`json:"image"`
}
func (g MigrateImageDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
