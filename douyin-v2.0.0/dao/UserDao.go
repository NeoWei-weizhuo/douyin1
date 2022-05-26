package dao

/*
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
}*/

//func (*UserDao) updateHeader(id int64, headerUrl string) bool {
//	return true
//}
//func (*UserDao) selectByEmail(email string) *User {
//	return nil
//}
