package handlers

import (
	"BackendAndroid/database"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetPoetry(c *gin.Context) {
	fetchedResponse, err := database.FetchRandomFromDB(database.DB, 1)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, fetchedResponse[0])

}

func GetPoetryList(c *gin.Context) {
	idStr := c.Param("n")
	n, err := strconv.Atoi(idStr)

	fetchedResponse, err := database.FetchRandomFromDB(database.DB, n)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, fetchedResponse)
}
