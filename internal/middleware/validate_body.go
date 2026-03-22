package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/natz/go-lib-app/internal/response"
)

func ValidateBody[T any]() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body T
		if err := c.ShouldBindJSON(&body); err != nil {
			log.Printf("Invalid request body: %s", err)
			response.BadRequest(c, `"Invalid request body"`+err.Error())
			c.Abort()
			return
		}
		c.Set("body", body)
		c.Next()
	}
}
