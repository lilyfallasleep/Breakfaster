package common

import (
	"github.com/gin-gonic/gin"
)

// ErrResponse is the error response type
type ErrResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

// Response is a helper function for sending error responses
func Response(c *gin.Context, httpCode, errCode int, err error) {
	message := err.Error()
	c.JSON(httpCode, ErrResponse{
		Code:    errCode,
		Message: message,
	})
}
