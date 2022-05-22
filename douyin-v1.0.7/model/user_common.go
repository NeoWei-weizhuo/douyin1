package model

import "time"

type User struct {
	Id            int64     `gorm:"column:user_id" json:"id,omitempty"`
	Username      string    `gorm:"column:user_name" json:"username,omitempty"`
	Password      string    `json:"password,omitempty"`
	Salt          string    `json:"salt,omitempty"`
	Followcount   int64     `gorm:"column:follow_count" json:"followcount,omitempty"`
	Followercount int64     `gorm:"column:follower_count" json:"followercount,omitempty"`
	CreateTime    time.Time `json:"createTime,omitempty"`
	IsFollow      bool   `gorm:"-" json:"is_follow,omitempty"`
}

func (du *User) TableName() string {
	return "user_t"
}

func (u *User) SetUsername(name string) {
	u.Username = name
}

func (u *User) SetPassword(password string) {
	u.Password = password
}

func (u *User) SetSalt(salt string) {
	u.Salt = salt
}

func (u *User) SetFollowcount(followcount int64) {
	u.Followcount = followcount
}

func (u *User) SetFollowercount(followercount int64) {
	u.Followercount = followercount
}

func (u *User) SetCreateTime(createTime time.Time) {
	u.CreateTime = createTime
}

type LoginTicket struct {
	Id      int64     `json:"id,omitempty"`
	UserId  int64     `json:"username,omitempty"`
	Ticket  string    `json:"password,omitempty"`
	Status  int       `json:"salt,omitempty"`
	Expired time.Time `json:"email,omitempty"`
}
//func (du *LoginTicket)TableName() (string) {
//	return "login_ticket"
//}

func (l *LoginTicket) SetUserId(userId int64) {
	l.UserId = userId
}
func (l *LoginTicket) SetTicket(ticket string) {
	l.Ticket = ticket
}
func (l *LoginTicket) SetStatus(status int) {
	l.Status = status
}
func (l *LoginTicket) SetExpired(expired time.Time) {
	l.Expired = expired
}
