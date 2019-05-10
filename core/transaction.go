package core

import "github.com/louisevanderlith/husk"

type Transaction struct {
	EntityKey    husk.Key
	Total        int64
	Requisitions []Requisition
}

func (o Transaction) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}

func (t Transaction) OpenRequsition(quoteKey husk.Key) {
	//trans, err :=
}

func GetTransactions(entityKey husk.Key) {

}
