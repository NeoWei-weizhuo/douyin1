package model

import "time"

type FollowT struct {
	Id         int64 `gorm:"primaryKey;autoIncrement;column:id" json:"id,omitempty"`
	UserId     int64 `gorm:"column:user_id" json:"user_id,omitempty"`
	ToUserId   int64 `gorm:"column:to_user_id" json:"to_user_id,omitempty"`
	FollowTime time.Time `gorm:"column:follow_time" json:"create_time,omitempty"`
}
func (du *FollowT)TableName() string {
	return "follow_t"
}

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

