package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Success   bool        `json:"success"`
	Error     string      `json:"error"`
	RequestID interface{} `json:"request_id,omitempty"`
}

func Error(c *gin.Context, status int, message string) {

	requestID, _ := c.Get("request_id")

	res := ErrorResponse{
		Success:   false,
		Error:     message,
		RequestID: requestID,
	}

	c.JSON(status, res)
}

func BadRequest(c *gin.Context, message string) {
	Error(c, http.StatusBadRequest, message)
}

func NotFound(c *gin.Context, message string) {
	Error(c, http.StatusNotFound, message)
}

func Internal(c *gin.Context) {
	Error(c, http.StatusInternalServerError, "internal server error")
}
