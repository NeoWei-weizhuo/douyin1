package dao

import (
	"douyin/conf"
	"douyin/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"sync"
)

type LoginTicketDao struct {
}

var (
	loginTicketDao  *LoginTicketDao
	loginTicketOnce sync.Once
)

func NewLoginTicketDaoInstance() *LoginTicketDao {
	loginTicketOnce.Do(
		func() {
			loginTicketDao = &LoginTicketDao{}
		})
	return loginTicketDao
}
func (*LoginTicketDao) InsertLoginTicket(loginTicket *model.LoginTicket) int64 {
	return InsertLoginTicket(loginTicket)
}

func (*LoginTicketDao) SelectByTicket(ticket string) *model.LoginTicket {
	return SelectByTicket(ticket)
}

func (*LoginTicketDao) UpdateLoginStatus(status int, ticket string) int {
	return UpdateLoginStatus(status, ticket)
}


func InsertLoginTicket(loginTick *model.LoginTicket) int64 {

	db := conf.DB
	db.Create(loginTick)
	return loginTick.Id

	//sqlStr := "insert into login_ticket (user_id,ticket,status,expired) values(?,?,?,?)"
	//ret, err := db.Exec(sqlStr, loginTick.UserId, loginTick.Ticket, loginTick.Status, loginTick.Expired)
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

func SelectByTicket(ticket string) *model.LoginTicket {
	//s := "select * from login_ticket where ticket = ?"
	var l model.LoginTicket
	db := conf.DB
	result := db.Find(&l, "ticket = ?", ticket)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Printf("scan failed, err :%v\n", result.Error)
		return nil
	}
	return &l

	//err := db.QueryRow(s, ticket).Scan(&l.Id, &l.UserId, &l.Ticket, &l.Status, &l.Expired)
	//if err != nil {
	//	fmt.Printf("scan failed, err :%v\n", err)
	//	return nil
	//}
	//return &l
}

func UpdateLoginStatus(status int, ticket string) int {
	db := conf.DB
	var l *model.LoginTicket

	err := db.Model(&l).Where("ticket = ?", ticket).Update("status", status)
	if err != nil {
		fmt.Printf("update failed, err: %v\n", err)
		return 0
	}
	return 1

	//s := "update login_ticket set status = ? where ticket = ?"
	//ret, err := db.Exec(s, status, ticket)
	//if err != nil {
	//	fmt.Printf("update LoginStatus failed, err: %v\n", err)
	//	return 0
	//}
	//rows, err := ret.RowsAffected()
	//if err != nil {
	//	fmt.Printf("update LoginStatus rows failed, err: %v\n", err)
	//	return 0
	//}
	//fmt.Printf("update LoginStatus success, update row: %v\n", rows)
	//return 1
}
