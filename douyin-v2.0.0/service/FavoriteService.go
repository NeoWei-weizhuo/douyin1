package service

import (
	"douyin/dao"
	"douyin/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func FavoriteAction(c *gin.Context, token string) model.Response {

	// 登陆鉴权复用RelationService.GetUserIdFromUserInfo
	uId, resp := GetUserIdFromUserInfo(token)

	if resp.StatusCode == -1 {
		return model.Response{StatusCode: -1, StatusMsg: "操作失败"}
	}

	actionType, err := strconv.ParseInt(c.Query("action_type"), 10, 32)
	if err != nil || (actionType != 1 && actionType != 2) {
		resp.StatusCode = -1
		resp.StatusMsg = "请求参数有误"
		return resp
	}
	videoId, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = "请求参数有误"
		return resp
	}

	if actionType == 2 {
		actionType = -1
	}

	success := dao.DoFavorite(uId, videoId, float32(actionType))
	if !success {
		resp.StatusCode = -1
		resp.StatusMsg = "操作失败"
		return resp
	}

	resp.StatusCode = 0
	return resp
}

func GetFavoriteVideos(c *gin.Context, token string) model.FavoriteListResponse {
	var videoList []model.Video
	userId, resp := GetUserIdFromUserInfo(token)
	if resp.StatusCode == -1 {
		return model.FavoriteListResponse{
			Response:  resp,
			VideoList: []model.Video{},
		}
	}

	videoList = dao.GetFavoriteList(userId)
	resp.StatusCode = 0
	return model.FavoriteListResponse{
		Response:  resp,
		VideoList: videoList,
	}
}
