package handlers

import (
	"fmt"
	"net/http"

	"BackendAndroid/database"
	"BackendAndroid/models"
	"BackendAndroid/utils"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		ErrorResponse(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Hash the password
	if err := user.HashPassword(); err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "Could not hash password")
		return
	}

	// Save the user to the database
	if err := models.CreateUser(database.DB, &user); err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "Could not create user")
		return
	}

	SuccessResponse(c, http.StatusCreated, "User registered successfully", nil)
}

func Login(c *gin.Context) {
	var inputUser models.User
	if err := c.ShouldBindJSON(&inputUser); err != nil {
		ErrorResponse(c, http.StatusBadRequest, "Invalid request payload")
		return
	}
	fmt.Println(inputUser)

	// Fetch the user from the database
	storedUser, err := models.GetUserByUsername(database.DB, inputUser.Username)
	if err != nil {
		ErrorResponse(c, http.StatusUnauthorized, "User not found")
		return
	}

	// Check the password
	if !storedUser.CheckPassword(inputUser.Password) {
		ErrorResponse(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	// Generate JWT Token
	token, err := utils.GenerateToken(storedUser.Username)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "Could not generate token")
		return
	}

	// Return token in the response
	SuccessResponse(c, http.StatusOK, "Login successful", gin.H{
		"username": storedUser.Username,
		"token":    token,
	})
}
