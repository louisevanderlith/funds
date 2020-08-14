package core

import (
	"github.com/louisevanderlith/husk"
)

type LineItem struct {
	StockKey       husk.Key
	Code           string
	Description    string
	UnitCost       int64
	UnitsRequested int
	UnitReceived   int
}

func (o LineItem) Valid() error {
	return husk.ValidateStruct(o)
}
