package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/natz/go-lib-app/internal/modules/rbac"
	"github.com/natz/go-lib-app/internal/response"
)

func RequirePermission(service *rbac.Service, permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleName := c.GetString("role")

		if !service.HasPermission(roleName, permission) {
			response.Forbidden(c, "You do not have permission to access this resource")
			c.Abort()
			return
		}

		c.Next()
	}
}
