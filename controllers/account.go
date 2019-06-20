package controllers

import (
	"github.com/louisevanderlith/mango/control"
)

type AccountController struct {
	control.APIController
}

func NewAccountCtrl(ctrlMap *control.ControllerMap) *AccountController {
	result := &AccountController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

// @Title GetCreditBalance
// @Description Gets the current user's credit balance.
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [get]
func (req *AccountController) Get() {

}
