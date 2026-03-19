package validators

import (
	"github.com/gin-gonic/gin"
)

func GetBody[T any](c *gin.Context) T {
	return c.MustGet("body").(T)
}
