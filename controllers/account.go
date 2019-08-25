package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

type AccountController struct {
}

// @Title GetCreditBalance
// @Description Gets the current user's credit balance.
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [get]
func (req *AccountController) Get(ctx context.Contexer) (int, interface{}) {
	return http.StatusNotImplemented, nil
}
