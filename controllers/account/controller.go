package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Title GetCreditBalance
// @Description Gets the current user's credit balance.
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [get]
func Get(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, nil)
}
