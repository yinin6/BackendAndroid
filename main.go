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

	router.GET("/poetry", handlers.GetPoetry)
	router.GET("/poetryList/:n", handlers.GetPoetryList)
	router.POST("/addFavorites", handlers.AddFavorite)
	router.POST("/removeFavorites", handlers.RemoveFavorites)
	router.GET("/getFavorites/:username", handlers.GetFavorites)
	router.GET("/getFavoritesList/:username", handlers.GetFavoritesList)

	router.POST("/saveNote", handlers.SaveNote)
	router.GET("/delNote/:id", handlers.DelNote)
	router.GET("/notes", handlers.GetNotes)
	router.GET("/getUserNotes/:username", handlers.GetNoteByUsername)

	// Protected routes
	authGroup := router.Group("/api")
	authGroup.Use(middleware.AuthMiddleware())
	{
		authGroup.GET("/profile", func(c *gin.Context) {
			username := c.MustGet("username").(string)
			c.JSON(http.StatusOK, gin.H{"message": "Welcome, " + username})
		})
	}
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
