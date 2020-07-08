package core

import (
	"github.com/louisevanderlith/funds/core/requisitionstatus"
	"github.com/louisevanderlith/husk"
)

type Requisition struct {
	QuoteKey    husk.Key
	Reference   string
	Status      requisitionstatus.Enum
	DebtorKey   husk.Key
	CreditorKey husk.Key
	Total       int64
	LineItems   []LineItem
}

func (o Requisition) Valid() error {
	return husk.ValidateStruct(&o)
}
