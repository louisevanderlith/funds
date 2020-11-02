package handles

import (
	"net/http"
)

// @Title GetCreditBalance
// @Description Gets the current user's credit balance.
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [get]
func GetAccounts(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "", http.StatusNotImplemented)
}
