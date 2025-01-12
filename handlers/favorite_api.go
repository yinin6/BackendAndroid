package handlers

import (
	"BackendAndroid/database"
	"BackendAndroid/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AddFavorite(c *gin.Context) {
	var request struct {
		UserID string `json:"user_id"`
		PoemID string `json:"poem_id"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		ErrorResponse(c, http.StatusBadRequest, "无效的请求")
		return
	}

	// 将收藏关系存入数据库或缓存
	err := database.AddToFavorites(request.UserID, request.PoemID)
	if err != nil {
		log.Println(err.Error(), request.UserID, request.PoemID)
		ErrorResponse(c, http.StatusInternalServerError, "收藏失败")
		return
	}

	SuccessResponse(c, http.StatusOK, "古诗已收藏", nil)
}

func RemoveFavorites(c *gin.Context) {
	var request struct {
		UserID string `json:"user_id"`
		PoemID string `json:"poem_id"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "无效的请求"})
		return
	}

	// 从数据库或缓存中移除收藏关系
	err := database.RemoveFromFavorites(request.UserID, request.PoemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "取消收藏失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "古诗已取消收藏"})
}

func GetFavorites(c *gin.Context) {
	userID := c.Query("username")

	// 从数据库或缓存中获取用户的收藏列表
	poemIDs, err := database.GetUserFavorites(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "获取收藏列表失败"})
		return
	}

	var poems []models.APIResponse

	for _, poemID := range poemIDs {
		// 根据 poemIDs 查询古诗详情
		poem, _ := database.FetchFromDB(poemID)
		poems = append(poems, *poem)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   poems,
	})
}

func GetFavoritesList(c *gin.Context) {
	userID := c.Param("username")

	log.Println("GetFavoritesList: ", userID)

	// 从数据库或缓存中获取用户的收藏列表
	poemIDs, err := database.GetUserFavorites(userID)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "获取收藏列表失败")
		return
	}
	SuccessResponse(c, http.StatusOK, "获取收藏列表成功", poemIDs)

}
