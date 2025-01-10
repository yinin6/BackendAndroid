package main

import (
	"BackendAndroid/database"
	"BackendAndroid/handlers"
	"BackendAndroid/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Initialize the database
	database.InitDB()

	// Create a Gin router
	router := gin.Default()

	// Public routes
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)

	// Protected routes
	authGroup := router.Group("/api")
	authGroup.Use(middleware.AuthMiddleware())
	{
		authGroup.GET("/profile", func(c *gin.Context) {
			username := c.MustGet("username").(string)
			c.JSON(http.StatusOK, gin.H{"message": "Welcome, " + username})
		})

	}

	// Start the server
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
