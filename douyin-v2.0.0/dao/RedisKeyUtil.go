package dao

func RedisKey(token string) string {

	const SPLIT = ":"
	const PREFIX_TICKET = "ticket"

	// 登录的凭证
	return PREFIX_TICKET + SPLIT + token

}
