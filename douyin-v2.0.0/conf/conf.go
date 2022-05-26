package conf

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
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

// InitRedis ...
func InitRedis(host string, passwd string) error {
	var pool = &redis.Pool{
		MaxIdle:     50,
		MaxActive:   100,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host)
			if err != nil {
				return nil, err
			}
			if passwd != "" {
				if _, err := c.Do("AUTH", passwd); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		// custom connection test method
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if _, err := c.Do("PING"); err != nil {
				return err
			}
			return nil
		},
	}
	RedisConn = &RedisConnc{pool}
	return nil
}
