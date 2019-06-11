package core

import "github.com/louisevanderlith/husk"

type Account struct {
	EntityKey    husk.Key
	Total        Total
	Requisitions []Requisition
}

func (o Account) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}

func (t Account) OpenRequsition(quoteKey husk.Key) {
	//trans, err :=
}

func GetTransactions(entityKey husk.Key) {

}
