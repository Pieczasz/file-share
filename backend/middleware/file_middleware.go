package middleware

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "path/filepath"
)

const MaxFileSize = 10 << 20 // 10MB

func ValidateFileUpload() gin.HandlerFunc {
    return func(c *gin.Context) {
        file, err := c.FormFile("file")
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
            c.Abort()
            return
        }

        if file.Size > MaxFileSize {
            c.JSON(http.StatusBadRequest, gin.H{"error": "File too large"})
            c.Abort()
            return
        }

        ext := filepath.Ext(file.Filename)
        allowedExts := map[string]bool{
            ".pdf":  true,
            ".doc":  true,
            ".docx": true,
            ".txt":  true,
            ".png":  true,
            ".jpg":  true,
            ".jpeg": true,
            ".gif":  true,
            ".zip":  true,
            ".tar":  true,
            ".gz":   true,
            ".mp4":  true,
            ".mp3":  true,
        }

        if !allowedExts[ext] {
            c.JSON(http.StatusBadRequest, gin.H{"error": "File type not allowed"})
            c.Abort()
            return
        }

        c.Next()
    }
}