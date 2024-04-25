package entity

import (
    "github.com/alfarady/letgo/lib"
)

type GetProductItemVariation2ResponseEntity struct{
    Name	string	`json:"name"`
    HasImage	bool	`json:"has_image"`
    Customize	bool	`json:"customize"`
    Options	[]string	`json:"options"`
}
func (g GetProductItemVariation2ResponseEntity) String() string {
    return lib.ObjectToString(g)
}
