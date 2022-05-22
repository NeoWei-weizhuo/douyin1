package model

import "time"

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

//type Video struct {
//	//Id            int64  `json:"id,omitempty"`
//	//Author        User   `json:"author"`
//	//PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
//	//CoverUrl      string `json:"cover_url,omitempty"`
//	//FavoriteCount int64  `json:"favorite_count,omitempty"`
//	//CommentCount  int64  `json:"comment_count,omitempty"`
//	//IsFavorite    bool   `json:"is_favorite,omitempty"`
//
//	Id            int64  `gorm:"column:id" json:"id,omitempty"`
//	AuthorId      int64  `gorm:"column:author_id" json:"author_id,omitempty"`
//	Author        User   `gorm:"foreignKey: Id;references: AuthorId" json:"author"` // gorm:"foreignKey: Id;reference:AuthorId"
//	PlayUrl       string `gorm:"column:play_url"  json:"play_url,omitempty"`
//	CoverUrl      string `gorm:"column:cover_url" json:"cover_url,omitempty"`
//	FavoriteCount int64  `gorm:"column:favorite_count" json:"favorite_count,omitempty"`
//	CommentCount  int64  `gorm:"column:comment_count" json:"comment_count,omitempty"`
//	IsFavorite    bool   `gorm:"-" json:"is_favorite,omitempty"`
//	CreateTime    time.Time
//}

func (du *Video) TableName() string {
	return "video_t"
}

type Video struct {
	Id            int64  `gorm:"column:id" json:"id,omitempty"`
	AuthorId      int64  `gorm:"column:author_id" json:"author_id,omitempty"`
	Author        User   `gorm:"foreignKey: user_id;references: AuthorId" json:"author"` // gorm:"foreignKey: Id;reference:AuthorId"
	PlayUrl       string `gorm:"column:play_url"  json:"play_url,omitempty"`
	CoverUrl      string `gorm:"column:cover_url" json:"cover_url,omitempty"`
	FavoriteCount int64  `gorm:"column:favorite_count" json:"favorite_count,omitempty"`
	CommentCount  int64  `gorm:"column:comment_count" json:"comment_count,omitempty"`
	IsFavorite    bool   `gorm:"-" json:"is_favorite,omitempty"`
	CreateTime    time.Time
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

//type User_t struct {
//	Id            int64  `gorm:"column:user_id" json:"id,omitempty"`
//	Name          string `gorm:"column:user_name" json:"name,omitempty"`
//	FollowCount   int64  `gorm:"column:follow_count" json:"follow_count,omitempty"`
//	FollowerCount int64  `gorm:"column:follower_count" json:"follower_count,omitempty"`
//	IsFollow      bool   `gorm:"-" json:"is_follow,omitempty"`
//}

type DatabaseVideo struct {
	Id            int64     `gorm:"primary_key"`
	AuthorId      int64     `gorm:"column:author_id"`
	PlayUrl       string    `gorm:"column:play_url"`
	CoverUrl      string    `gorm:"column:cover_url"`
	FavoriteCount int64     `gorm:"column:favorite_count"`
	CommentCount  int64     `gorm:"comment_count"`
	Title         string    `gorm:"column:title"`
	Timestamp     time.Time `gorm:"column:create_time"`
}

func (dv DatabaseVideo) TableName() string {
	return "video_t"
}

type DatabaseUser struct {
	Id            int64  `gorm:"primary_key;column:user_id"`
	Name          string `gorm:"column:user_name"`
	Password      string `gorm:"column:password"`
	FollowCount   int64  `gorm:"column:follow_count"`
	FollowerCount int64  `gorm:"column:follower_count"`
}

func (du DatabaseUser) TableName() string {
	return "user_t"
}
