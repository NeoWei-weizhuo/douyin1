package dao

import (
	"time"

	"douyin/conf"
	"douyin/model"
	"gorm.io/gorm"
)

func DoFollow(userId, toUserId int64) bool {
	var err error
	db := conf.DB
	var user  *model.User

	err = db.Model(&user).Where("user_id = ? ", userId).Update("follow_count", gorm.Expr("follow_count + ?", 1)).Error
	if err != nil {
		return false
	}

	//user_t.toUser粉丝数+1
	err = db.Model(&user).Where("user_id = ? ", toUserId).Update("follower_count", gorm.Expr("follower_count + ?", 1)).Error
	if err != nil {
		return false
	}

	//将toUser加入User的关注列表——follow_t
	db.Create(&model.FollowT{
		UserId:     userId,
		ToUserId:   toUserId,
		FollowTime: time.Now(),
	})

	//将User加入toUser的粉丝列表——暂时没有用粉丝表
	//db.Create(&model.FollowT{
	//	UserId:     toUserId,
	//	ToUserId:   userId,
	//	FollowTime: time.Now().Unix(),
	//})

	return true
}

func CancelFollow(userId, toUserId int64) bool {
	var err error
	db := conf.DB
	var user  *model.User
	//User关注数-1
	err = db.Model(&user).Where("user_id = ? ", userId).Update("follow_count", gorm.Expr("follow_count - ?", 1)).Error
	if err != nil {
		return false
	}


	//toUser粉丝数-1
	err = db.Model(&user).Where("user_id = ? ", toUserId).Update("follower_count", gorm.Expr("follow_count - ?", 1)).Error
	if err != nil {
		return false
	}

	var follow model.FollowT
	//将toUser移出User的关注列表
	db.Where("user_id = ?", userId).Where("to_user_id = ?", toUserId).Delete(&follow)

	//将User移出toUser的粉丝列表——暂无粉丝表
	//conf.DB.Where("user_id = ?", toUserId).Where("to_user_id = ?", userId).First(&follow)
	//conf.DB.Delete(&follow)

	return true
}

func GetUserListById(id int64, flag int) []model.User {
	var userList []model.User
	var followRows []model.FollowT
	var db = conf.DB

	//在关注表中，以user_id去查找时，获取到的记录对应的to_user_id就是被关注者的user_id
	//从而由该记录找到to_user_id，以它作为user_id获取到被关注者的用户列表
	//同理，以to_user_id去查找时，获取到的记录对应的user_id就是被粉丝的user_id
	//以它作为user_id获取到粉丝的用户列表
	if flag == 1 {
		//查找我关注的人
		//select to_user_id from follow_t where user_id=id;
		db.Where("user_id = ?", id).Find(&followRows)
		//遍历构建userList
		for _, v := range followRows {
			user := model.User{
				Id: v.Id,
			}
			userList = append(userList, user)
		}
	}else if flag == 2 {
		//查找我的粉丝，也即关注我的人
		//select user_id from follow_t where to_user_id=id;
		db.Where("to_user_id = ?", id).Find(&followRows)
		//遍历构建userList
		for _, v := range followRows {
			user := model.User{
				Id: v.Id,
			}
			userList = append(userList, user)
		}
	}



	return userList
}
