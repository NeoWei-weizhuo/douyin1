package controller

import (
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	c.JSON(http.StatusOK, service.FavoriteAction(c, token))
}

func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	c.JSON(http.StatusOK, service.GetFavoriteVideos(c, token))
}

func CommentAction(c *gin.Context) {
	token := c.Query("token")
	c.JSON(http.StatusOK, service.CommentAction(c, token))
}

func CommentList(c *gin.Context) {
	token := c.Query("token")
	c.JSON(http.StatusOK, service.GetComments(c, token))
}
