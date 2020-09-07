package core

import (
	"github.com/louisevanderlith/funds/core/requisitionstatus"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/validation"
)

type Requisition struct {
	QuoteKey    hsk.Key
	Reference   string
	Status      requisitionstatus.Enum
	DebtorKey   hsk.Key //Hero
	CreditorKey hsk.Key //Hero
	Total       int64
	LineItems   []LineItem
}

func (o Requisition) Valid() error {
	return validation.Struct(o)
}
