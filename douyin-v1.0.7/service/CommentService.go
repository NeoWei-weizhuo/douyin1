package service

import (
	"douyin/dao"
	"douyin/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CommentAction(c *gin.Context, token string) model.Response {

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

	commentText := c.Query("comment_text")

	//对删除评论id单独删除的操作

	if actionType == 2 {
		actionType = -1
	}

	if !dao.DoComment(uId, videoId, commentText, float32(actionType)) {
		resp.StatusCode = -1
		resp.StatusMsg = "操作失败"
		return resp
	}

	resp.StatusCode = 0
	return resp
}

func GetComments(c *gin.Context, token string) model.CommentListResponse {
	var commentList []model.Comment
	_, resp := GetUserIdFromUserInfo(token)
	if resp.StatusCode == -1 {
		return model.CommentListResponse{
			Response:    resp,
			CommentList: []model.Comment{},
		}
	}
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)

	commentList = dao.GetCommentsList(videoId)
	resp.StatusCode = 0
	return model.CommentListResponse{
		Response:    resp,
		CommentList: commentList,
	}
}
