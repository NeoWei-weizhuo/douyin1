package controller

import (
	"douyin/model"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedResponse struct {
	model.Response
	VideoList []model.Video `json:"video_list,omitempty"`
	NextTime  int64           `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	// 最新时间
	latestTime := c.DefaultQuery("latest_time", "")
	// 用户token
	token := c.DefaultQuery("token", "")

	c.JSON(http.StatusOK, FeedResponse{
		Response:  model.Response{StatusCode: 0},
		VideoList: service.GetListByLatestTimeAndToken(token, latestTime, 30),
		NextTime:  time.Now().Unix(),
	})
}
