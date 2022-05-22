package dao

//type LoginTicketDao struct {
//}
//
//var (
//	loginTicketDao  *LoginTicketDao
//	loginTicketOnce sync.Once
//)
//
//func NewLoginTicketDaoInstance() *LoginTicketDao {
//	loginTicketOnce.Do(
//		func() {
//			loginTicketDao = &LoginTicketDao{}
//		})
//	return loginTicketDao
//}
//func (*LoginTicketDao) InsertLoginTicket(loginTicket *model.LoginTicket) int64 {
//	return InsertLoginTicket(loginTicket)
//}
//
//func (*LoginTicketDao) SelectByTicket(ticket string) *model.LoginTicket {
//	return SelectByTicket(ticket)
//}
//
//func (*LoginTicketDao) UpdateLoginStatus(status int, ticket string) int {
//	return UpdateLoginStatus(status, ticket)
//}
