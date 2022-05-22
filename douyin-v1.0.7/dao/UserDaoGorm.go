package dao

import (
	"douyin/conf"
	"douyin/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"sync"
)

type UserDao struct {
}

var (
	userDao  *UserDao
	userOnce sync.Once
)

func NewUserDaoInstance() *UserDao {
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		})
	return userDao
}
func (*UserDao) SelectById(id int64) *model.User {
	return QueryById(id)
}

func (*UserDao) SelectByName(username string) *model.User {
	return QueryByName(username)
}

func (*UserDao) InsertUser(user *model.User) int64 {
	return InsertUser(user)
}

func (*UserDao) UpdatePassword(id int64, password string) int {
	return UpdatePassword(id, password)
}

func QueryById(id int64) *model.User {

	var u model.User
	db := conf.DB
	result := db.Find(&u, "user_id = ?", id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &u
	}
	return &u

}

func QueryByName(name string) *model.User {

	var u model.User
	db := conf.DB
	result := db.First(&u, "user_name = ?", name)
	//db.Find(&u, "user_name = ?", name)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &u

}

func UpdatePassword(id int64, password string) int {
	db := conf.DB
	var user *model.User

	err := db.Model(&user).Where("user_id = ?", id).Update("password",password)
	if err != nil {
		fmt.Printf("update failed, err: %v\n", err)
		return 0
	}
	return 1
	//s := "update user set password = ? where id = ?"
	//ret, err := db.Exec(s, password, id)
	//if err != nil {
	//	fmt.Printf("update failed, err: %v\n", err)
	//	return 0
	//}
	//rows, err := ret.RowsAffected()
	//if err != nil {
	//	fmt.Printf("update rows failed, err: %v\n", err)
	//	return 0
	//}
	//fmt.Printf("update success, update row: %v\n", rows)
	//return 1
}

func InsertUser(user *model.User) int64 {

	db := conf.DB
	db.Create(user)
	return user.Id
	//sqlStr := "insert into user_t (user_name,password,follow_count,follower_count,salt,create_time) " +
	//	"values(?,?,?,?,?,?)"
	//ret, err := db.Exec(sqlStr, user.Username, user.Password, user.Followcount,
	//	user.Followercount, user.Salt, user.CreateTime)
	//if err != nil {
	//	fmt.Printf("insert data error: %v\n", err)
	//	return 0
	//}
	//theID, err := ret.LastInsertId()
	//if err != nil {
	//	fmt.Printf("get LastInsertId error: %v\n", err)
	//	return 0
	//}
	//fmt.Printf("insert success, the id: %v\n", theID)
	//return theID
}