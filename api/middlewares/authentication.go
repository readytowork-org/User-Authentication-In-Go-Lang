package middlewares

import (
	"digitalsign-api/api/responses"
	"digitalsign-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetMiddlewareAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := utils.TokenValid(c)
		if err != nil {
			responses.ErrorJSON(c, http.StatusForbidden, "Unauthorized")
			c.Abort()
			return
		}

	}
}
