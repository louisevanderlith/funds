package core

import (
	"github.com/louisevanderlith/husk"
)

type context struct {
	Transactions husk.Tabler
}

var ctx context

func CreateContext() {
	ctx = context{
		Transactions: husk.NewTable(new(Transaction)),
	}
}

func Shutdown() {
	ctx.Transactions.Save()
}
