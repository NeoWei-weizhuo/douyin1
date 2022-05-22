package dao

import (
	"douyin/conf"
	"douyin/model"
)

// GetListByLatestTime GetList 获取列表
func GetListByLatestTime(latestTime string, limit int) []model.Video {
	var VideoList []model.Video
	videoQuery := conf.DB.Model(model.Video{}).Joins("Author")
	// 存在latest_time 就走以下查询
	//if latestTime != "" {
	//	videoQuery.Where("video_t.create_time <= ?", latestTime)
	//}
	videoQuery.Order("create_time DESC").Limit(limit).Find(&VideoList)

	return VideoList
}
