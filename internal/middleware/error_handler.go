package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	appErrors "github.com/natz/go-lib-app/internal/errors"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err

		switch e := err.(type) {
		case *appErrors.AppError:

			c.JSON(e.StatusCode, gin.H{
				"success": false,
				"error":   e.Message,
			})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
	}
}
