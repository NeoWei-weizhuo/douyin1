package dao

import (
	"douyin/conf"
	"douyin/model"
	"errors"
	"gorm.io/gorm"
	"time"
)

type favor struct {
	Id           int64 `gorm:"primary_key"`
	VideoId      int64
	UserId       int64
	FavoriteTime time.Time
}

func (f favor) TableName() string {
	return "favorite_video_t"
}

// SaveVideoMessage 将视频信息存入Video表
func SaveVideoMessage(video model.DatabaseVideo) error {
	db := conf.DB
	if err := db.Create(&video).Error; err != nil {
		return err
	}
	return nil
}

// GetVideosByUserId 通过userId在Video表中查询用户上传的视频列表，按时间逆序返回
func GetVideosByUserId(id int64) ([]model.DatabaseVideo, error) {
	var videos []model.DatabaseVideo
	db := conf.DB
	result := db.Order("create_time desc").Find(&videos, "author_id = ?", id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return videos, nil
	}
	return videos, nil
}

func GetAuthorById(id int64) (model.DatabaseUser, error) {
	var author model.DatabaseUser
	err := conf.DB.First(&author, "user_id = ?", id).Error
	if err != nil {
		return author, err
	}
	return author, nil
}

func IsFavor(videoId int64, userId int64) (bool, error) {
	var f favor
	result := conf.DB.First(&f, "video_id = ? AND user_id = ?", videoId, userId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return true, nil
}
