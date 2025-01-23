package middleware

import (
    "github.com/gin-gonic/gin"
    "log"
    "time"
)

func Logger() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        path := c.Request.URL.Path

        c.Next()

        latency := time.Since(start)
        log.Printf("[%s] %s %s %v", c.Request.Method, path, c.ClientIP(), latency)
    }
}