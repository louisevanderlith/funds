package core

import (
	"github.com/louisevanderlith/husk"
)

type context struct {
	Accounts husk.Table
}

var ctx context

func CreateContext() {
	ctx = context{
		Accounts: husk.NewTable(Account{}),
	}
}

func Shutdown() {
	ctx.Accounts.Save()
}
