package dao

import (
	_ "github.com/go-sql-driver/mysql"
)

//var db *sql.DB
//
//func InitDB() (err error) {
//	dsn := "sql01:1234Bt!-@tcp(101.43.179.86:3306)/tt?charset=utf8mb4&parseTime=True"
//	db, err = sql.Open("mysql", dsn)
//	if err != nil {
//		return err
//	}
//	err = db.Ping()
//	if err != nil {
//		return err
//	}
//	return nil
//}

//func QueryById(id int64) *model.User {
//	s := "select * from user_t where user_id = ?"
//	var u model.User
//	err := db.QueryRow(s, id).Scan(&u.Id, &u.Username, &u.Password,
//		&u.Followcount, &u.Followercount, &u.Salt, &u.CreateTime)
//	if err != nil {
//		fmt.Printf("scan failed, err :%v\n", err)
//		return nil
//	}
//	fmt.Printf("name: %s\n", u.Username)
//	return &u
//}
//
//func QueryByName(name string) *model.User {
//	s := "select * from user_t where user_name = ?"
//	var u model.User
//	err := db.QueryRow(s, name).Scan(&u.Id, &u.Username, &u.Password,
//		&u.Followcount, &u.Followercount, &u.Salt, &u.CreateTime)
//	if err != nil {
//		fmt.Printf("scan failed, err :%v\n", err)
//		return nil
//	}
//	fmt.Printf("name: %s\n", u.Username)
//	return &u
//}
//
//func UpdatePassword(id int64, password string) int {
//	s := "update user set password = ? where id = ?"
//	ret, err := db.Exec(s, password, id)
//	if err != nil {
//		fmt.Printf("update failed, err: %v\n", err)
//		return 0
//	}
//	rows, err := ret.RowsAffected()
//	if err != nil {
//		fmt.Printf("update rows failed, err: %v\n", err)
//		return 0
//	}
//	fmt.Printf("update success, update row: %v\n", rows)
//	return 1
//}
//
//func InsertUser(user *model.User) int64 {
//	sqlStr := "insert into user_t (user_name,password,follow_count,follower_count,salt,create_time) " +
//		"values(?,?,?,?,?,?)"
//	ret, err := db.Exec(sqlStr, user.Username, user.Password, user.Followcount,
//		user.Followercount, user.Salt, user.CreateTime)
//	if err != nil {
//		fmt.Printf("insert data error: %v\n", err)
//		return 0
//	}
//	theID, err := ret.LastInsertId()
//	if err != nil {
//		fmt.Printf("get LastInsertId error: %v\n", err)
//		return 0
//	}
//	fmt.Printf("insert success, the id: %v\n", theID)
//	return theID
//}
