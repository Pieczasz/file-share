package handlers

import (
	"github.com/gin-gonic/gin"
)

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func HandleError(c *gin.Context, status int, message string) {
	c.JSON(status, APIError{
		Code:    status,
		Message: message,
	})
}