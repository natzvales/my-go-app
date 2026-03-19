package validators

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUUIDParam(c *gin.Context, key string) uuid.UUID {
	return c.MustGet(key).(uuid.UUID)
}
