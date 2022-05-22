package model

import "time"

type FavoriteVideoT struct {
	Id           int64     `gorm:"column:id"`
	UserId       int64     `gorm:"column:user_id"`
	VideoId      int64     `gorm:"column:video_id"`
	FavoriteTime time.Time `gorm:"column:favorite_time"`
}

type FavoriteListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

type CommentVideoT struct {
	Id          int64     `gorm:"column:id"`
	UserId      int64     `gorm:"column:user_id"`
	VideoId     int64     `gorm:"column:video_id"`
	CommentTime time.Time `gorm:"column:comment_time"`
	CommentText string    `gorm:"column:comment_text"`
}

type CommentListResponse struct {
	Response
	CommentList []Comment `gorm:"column:comment_list"`
}
