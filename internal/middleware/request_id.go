package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestID() gin.HandlerFunc {

	return func(c *gin.Context) {

		requestID := uuid.New().String()

		// store in gin context
		c.Set("request_id", requestID)

		// return in response header
		c.Writer.Header().Set("X-Request-ID", requestID)

		c.Next()
	}
}
