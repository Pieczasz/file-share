package main

import (
	"backend/database"
	"backend/routes"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
    "github.com/joho/godotenv"
    "log"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    database.InitDatabase()
    
    router := gin.Default()
    
    // Add CORS middleware
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"},
        AllowMethods:     []string{"GET", "POST", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

    routes.RegisterFileRoutes(router)

    router.Run(":8080")
}