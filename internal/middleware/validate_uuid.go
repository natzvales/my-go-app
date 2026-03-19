package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/natz/go-lib-app/internal/response"
)

func ValidateUUIDParam(key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		param := c.Param(key)
		parsed, err := uuid.Parse(param)
		if err != nil {
			response.BadRequest(c, "Invalid UUID parameter")
			c.Abort()
			return
		}
		c.Set(param, parsed)
		c.Next()
	}
}
