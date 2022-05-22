package dao

import (
	"douyin/conf"
	"douyin/model"
	"gorm.io/gorm"
	"math"
	"time"
)

type videoT struct {
}

// DoFavorite adder: 1 like; -1 cancel like
func DoFavorite(userId, videoId int64, adder float32) bool {
	if math.Abs(float64(adder)) != 1 {
		return false
	}
	// update video table for favorite_count
	db := conf.DB
	var video *model.Video
	var favor *model.FavoriteVideoT
	err := db.Model(&video).Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count + ?", adder)).Error
	if err != nil {
		return false
	}

	// update favorite_video_t
	if adder == 1 {
		err = db.Create(&model.FavoriteVideoT{
			UserId:       userId,
			VideoId:      videoId,
			FavoriteTime: time.Now(),
		}).Error
	} else {
		err = db.Where("user_id = ?", userId).Where("video_id = ?", videoId).Delete(&favor).Error
	}
	return err == nil
}

// GetFavoriteList query video_id by user_id and then query videos by video_id
func GetFavoriteList(userId int64) []model.Video {
	db := conf.DB
	var favorList []model.FavoriteVideoT
	var videoList []model.Video

	db.Where("user_id = ?", userId).Find(&favorList)
	var ids []int64
	for _, v := range favorList {
		ids = append(ids, v.VideoId)
	}

	db.Table("video_t").Where("id in (?)", ids).Find(&videoList)
	return videoList
}
