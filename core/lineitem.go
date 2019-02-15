package funds

import (
	"github.com/louisevanderlith/mango"
)

type LineItem struct {
	Description    string
	UnitCost       int64
	UnitsRequisted int
	UnitReceived   int
}

func (o LineItem) Valid() (bool, error) {
	return mango.ValidateStruct(&o)
}
