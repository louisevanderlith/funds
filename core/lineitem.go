package core

import (
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/validation"
)

type LineItem struct {
	StockKey       hsk.Key
	Code           string
	Description    string
	UnitCost       int64
	UnitsRequested int
	UnitReceived   int
}

func (o LineItem) Valid() error {
	return validation.Struct(o)
}
