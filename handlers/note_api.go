package handlers

import (
	"BackendAndroid/database"
	"BackendAndroid/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SaveNote(c *gin.Context) {
	var note models.Note
	if err := c.ShouldBindJSON(&note); err != nil {
		ErrorResponse(c, http.StatusBadRequest, "invalid note")
		return
	}

	// 保存 Note 到数据库
	err := database.SaveNote(database.DB, note)
	if err != nil {
		fmt.Println(err)
		return
	}
	SuccessResponse(c, http.StatusOK, "note saved successfully", nil)
}

func DelNote(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	err = database.DeleteNoteByID(id)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "failed to delete note")
		return
	}

	SuccessResponse(c, http.StatusOK, "Note deleted successfully", nil)

}

func GetNotes(c *gin.Context) {
	notes, err := database.GetNotes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	SuccessResponse(c, http.StatusOK, "all Notes fetched successfully", notes)
}

func GetNoteByUsername(c *gin.Context) {
	username := c.Param("username")
	notes, err := database.GetNotesByUsername(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	SuccessResponse(c, http.StatusOK, "user's Notes fetched successfully", notes)

}
