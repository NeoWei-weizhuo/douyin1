package conf

import (
	"fmt"
)

// HostAndPort 获得主机ip和端口，供存储视频的url使用
func HostAndPort() string {
	return "47.98.196.156:8080"
}

// Init 初始化配置项
func Init() {
	//连接数据库
	MysqlDsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=%s&timeout=%s",
		"sql01",    // 用户名
		"1234Bt!-", // 密码
		"101.43.179.86",
		"3306",
		"tt",
		"Asia%2FShanghai",
		"10s",
	)

	Database(MysqlDsn)
}
