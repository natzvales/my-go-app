package middleware

import (
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/natz/go-lib-app/internal/response"
)

func RequireRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {

		user := GetUser(c)
		roleValue, exists := user.Role, true

		if !exists {
			response.Forbidden(c, "Forbidden")
			c.Abort()
			return
		}

		role := roleValue

		if slices.Contains(roles, role) {
			c.Next()
			return
		}

		response.Forbidden(c, "Insufficient permissions")
		c.Abort()

	}
}
