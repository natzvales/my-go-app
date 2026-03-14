package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Success   bool        `json:"success"`
	Data      interface{} `json:"data,omitempty"`
	Message   string      `json:"message,omitempty"`
	RequestID interface{} `json:"request_id,omitempty"`
}

func Success(c *gin.Context, data interface{}) {

	requestID, _ := c.Get("request_id")

	res := SuccessResponse{
		Success:   true,
		Data:      data,
		RequestID: requestID,
	}

	c.JSON(http.StatusOK, res)
}

func Created(c *gin.Context, data interface{}) {

	requestID, _ := c.Get("request_id")

	res := SuccessResponse{
		Success:   true,
		Data:      data,
		RequestID: requestID,
	}

	c.JSON(http.StatusCreated, res)
}
