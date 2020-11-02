package core

import (
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/validation"
)

type Account struct {
	HeroKey      hsk.Key
	Total        Total
	Requisitions []Requisition
}

func (o Account) Valid() error {
	return validation.Struct(o)
}

func (t Account) OpenRequsition(quoteKey hsk.Key) {
	//trans, err :=
}

func GetTransactions(entityKey hsk.Key) {

}
