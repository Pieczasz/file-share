package routes

import (
	"backend/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterFileRoutes(router *gin.Engine) {
	router.POST("/upload", handlers.UploadFile)
	router.GET("/download/:id", handlers.GetFile)
}