package core

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/husk/serials"
)

type context struct {
	Accounts husk.Tabler
}

var ctx context

func CreateContext() {
	ctx = context{
		Accounts: husk.NewTable(Account{}, serials.GobSerial{}),
	}
}

func Shutdown() {
	ctx.Accounts.Save()
}
