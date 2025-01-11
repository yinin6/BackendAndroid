package handlers

import (
	"BackendAndroid/database"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetPoetry(c *gin.Context) {
	fetchedResponse, err := database.FetchRandomFromDB(database.DB, 1)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, fetchedResponse[0])

}
