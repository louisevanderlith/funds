package core

import (
	"github.com/louisevanderlith/husk"
)

type LineItem struct {
	Description    string
	UnitCost       int64
	UnitsRequisted int
	UnitReceived   int
}

func (o LineItem) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}
