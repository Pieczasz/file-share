package main

import (
    "backend/database"
    "backend/routes"
    "github.com/gin-contrib/cors"
    "github.com/gin-contrib/timeout"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "log"
    "time"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    database.InitDatabase()
    
    router := gin.Default()

    // Middleware setup
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"},
        AllowMethods:     []string{"GET", "POST", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))

    // Simple rate limiting middleware
    router.Use(func(c *gin.Context) {
        // Add basic rate limiting logic here
        c.Next()
    })

    // Timeout middleware
    router.Use(timeout.New(
        timeout.WithTimeout(10*time.Second),
        timeout.WithHandler(func(c *gin.Context) {
            c.Next()
        }),
    ))

    routes.RegisterFileRoutes(router)

    router.Run(":8080")
}