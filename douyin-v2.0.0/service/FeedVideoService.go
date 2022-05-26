package service

import (
	"douyin/dao"
	"douyin/model"
	"fmt"
)

// GetListByLatestTimeAndToken 获取列表
func GetListByLatestTimeAndToken(token string, latestTime string, limit int) []model.Video {
	// 视频列表
	VideoList := dao.GetListByLatestTime(latestTime, limit)
	// token为空，直接返回视频列表
	if token == "" {
		return VideoList
	}
	// 应使用实际上的鉴权模式来鉴权，这里先模拟为false
	exist := false
	// 不存在该用户直接返回列表
	if !exist {
		return VideoList
	}

	// 存在该用户就遍历视频列表给isFavorite赋值
	for index, value := range VideoList {
		var likeExist int64
		fmt.Println(value)
		// 这里查询favorite_video_t表格, 示例： 假设favorite_video_t表格的结构体是VideoLike 则conf.DB.Model(VideoLike{}).Where("user_id <= ? AND video_id = ?", value.AuthorId,value.Id).limit(1).Count(&likeExist)
		// 假设likeExist > 0 则设置true
		VideoList[index].IsFavorite = likeExist > 0
	}

	return VideoList
}
