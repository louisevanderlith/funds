package core

import (
	"github.com/louisevanderlith/husk"
)

type LineItem struct {
	Stock
	Code string
	Description    string
	UnitCost       int64
	UnitsRequisted int
	UnitReceived   int
}

func (o LineItem) Valid() error {
	return husk.ValidateStruct(o)
}
