package routers

import (
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"
	"github.com/louisevanderlith/droxolite/routing"
	"github.com/louisevanderlith/funds/controllers"
)

func Setup(poxy resins.Epoxi) {
	//Credit
	credCtrl := &controllers.CreditController{}
	msgGroup := routing.NewRouteGroup("credit", mix.JSON)
	//msgGroup.AddRoute("/", "POST", roletype.Unknown, credCtrl.Post)
	msgGroup.AddRoute("All Credits", "/all/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, credCtrl.Get)
	//msgGroup.AddRoute("/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Unknown, credCtrl.GetOne)
	poxy.AddGroup(msgGroup)

	//Requisition
	reqCtrl := &controllers.RequisitionController{}
	reqGroup := routing.NewRouteGroup("requisition", mix.JSON)
	reqGroup.AddRoute("Create Requisition", "", "POST", roletype.Unknown, reqCtrl.Post)
	reqGroup.AddRoute("All Requistions", "/all/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, reqCtrl.Get)
	//reqGroup.AddRoute("/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Unknown, reqCtrl.GetOne)
	poxy.AddGroup(reqGroup)

	//Account
	accCtrl := &controllers.AccountController{}
	accGroup := routing.NewRouteGroup("account", mix.JSON)
	//accGroup.AddRoute("/", "POST", roletype.Unknown, accCtrl.Post)
	accGroup.AddRoute("All Accounts", "/all/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, accCtrl.Get)
	//accGroup.AddRoute("/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Unknown, accCtrl.GetOne)
	poxy.AddGroup(accGroup)
	/*ctrlmap := EnableFilter(s, host)

	beego.Router("/v1/credit", controllers.NewCreditCtrl(ctrlmap))
	beego.Router("/v1/requisition", controllers.NewRequisitionCtrl(ctrlmap))
	beego.Router("/v1/account", controllers.NewAccountCtrl(ctrlmap))*/
}

/*
func EnableFilter(s *mango.Service, host string) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(secure.ActionMap)
	emptyMap["GET"] = roletype.User
	emptyMap["POST"] = roletype.User
	emptyMap["PUT"] = roletype.User

	ctrlmap.Add("/v1/credit", emptyMap)
	ctrlmap.Add("/v1/requisition", emptyMap)
	ctrlmap.Add("/v1/account", emptyMap)

	beego.InsertFilter("/*", beego.BeforeRouter, ctrlmap.FilterAPI, false)
	allowed := fmt.Sprintf("https://*%s", strings.TrimSuffix(host, "/"))

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{allowed},
		AllowMethods: []string{"GET", "PUT", "POST", "OPTIONS"},
	}), false)

	return ctrlmap
}
*/
