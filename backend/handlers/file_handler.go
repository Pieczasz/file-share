package handlers

import (
	"backend/database"
	"backend/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const uploadDir = "/uploads"

func init() {
    if err := os.MkdirAll(uploadDir, 0755); err != nil {
        log.Fatal("couldn't create upload directory:", err)
    }
}

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fileID := uuid.New().String()
    savePath := filepath.Join(uploadDir, fileID + "_" + file.Filename)
	err = c.SaveUploadedFile(file, savePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fileModel := models.File{
		ID:        fileID,
		Name:      file.Filename,
		Extension: filepath.Ext(file.Filename),
		Path:      savePath,
		Size:      file.Size,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ExpiresAt: time.Now().Add(time.Hour * 24),
	}
	database.DB.Create(&fileModel)

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "file": fileModel})
}

func GetFile(c *gin.Context) {
    fileID := c.Param("id")
    var file models.File
    
    if err := database.DB.First(&file, "id = ?", fileID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
        return
    }
    
    if file.ExpiresAt.Before(time.Now()) {
        database.DB.Delete(&file)
        c.JSON(http.StatusGone, gin.H{"error": "File has expired"})
        return
    }
    
    c.File(file.Path)
}