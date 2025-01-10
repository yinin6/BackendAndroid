package handlers

import (
	"BackendAndroid/models"
	"github.com/gin-gonic/gin"
	"time"
)

func SuccessResponse(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(200, models.NetworkResponse{
		Status:    status,
		Message:   message,
		Data:      data,
		Timestamp: time.Now().Unix(),
	})
}

func ErrorResponse(c *gin.Context, status int, message string) {
	c.JSON(200, models.NetworkResponse{
		Status:    status,
		Message:   message,
		Data:      nil,
		Timestamp: time.Now().Unix(),
	})
}
