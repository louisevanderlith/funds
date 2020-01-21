package routers

import (
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"
	"github.com/louisevanderlith/funds/controllers"
)

func Setup(e resins.Epoxi) {
	credCtrl := &controllers.Credits{}
	reqCtrl := &controllers.Requisitions{}
	accCtrl := &controllers.Accounts{}
	e.JoinBundle("/", roletype.Owner, mix.JSON, credCtrl, reqCtrl, accCtrl)
}
