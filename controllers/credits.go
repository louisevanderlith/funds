package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

type Credits struct {
}

// @Title GetCreditBalance
// @Description Gets the current user's credit balance.
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [get]
func (req *Credits) Get(ctx context.Requester) (int, interface{}) {
	return http.StatusNotImplemented, nil
}
