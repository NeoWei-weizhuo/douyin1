package conf

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

// RedisConn redis链接单例
var RedisConn *RedisConnc

// RedisConnc ...
type RedisConnc struct {
	Pool *redis.Pool
}

func Redis(host string, password string) {

	err := InitRedis(host, password)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("redis连接成功!")
}
