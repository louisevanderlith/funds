package controllers

import "github.com/louisevanderlith/droxolite/xontrols"

type AccountController struct {
	xontrols.APICtrl
}

// @Title GetCreditBalance
// @Description Gets the current user's credit balance.
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [get]
func (req *AccountController) Get() {

}
