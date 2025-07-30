package response

import (
	"github.com/gin-gonic/gin"
)

// Response is a generic response structure.
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func ResponseSuccess(c *gin.Context, message string, data interface{}) {
	c.JSON(200, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func ReponseError(c *gin.Context, errCode int, message string) {
	c.JSON(errCode, Response{
		Success: false,
		Message: message,
	})
}
