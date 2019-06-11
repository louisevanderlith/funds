package core

import (
	"github.com/louisevanderlith/husk"
)

type context struct {
	Accounts husk.Tabler
}

var ctx context

func CreateContext() {
	ctx = context{
		Accounts: husk.NewTable(new(Account)),
	}
}

func Shutdown() {
	ctx.Accounts.Save()
}
