package dao

import (
	"crypto/md5"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

func GetRandomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var result []byte
	for i := 0; i < n; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

func GetRandomInt(n int) int {
	maxNum := n
	rand.Seed(time.Now().UnixNano())
	randNumber := rand.Intn(maxNum)
	return randNumber
}

func Md5(s string) string {

	// 进行md5加密，因为Sum函数接受的是字节数组，因此需要注意类型转换
	srcCode := md5.Sum([]byte(s))

	// md5.Sum函数加密后返回的是字节数组，需要转换成16进制形式
	code := fmt.Sprintf("%x", srcCode)

	return string(code)
}

func IsBlank(cs string) bool {
	strLen := len(cs)
	if cs != "" && strLen != 0 {
		for i := 0; i < strLen; i++ {
			if cs[i] == ' ' {
				return true
			}
		}
		return false
	} else {
		return true
	}
}

func CheckPassword(password string) (b bool) {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", password); !ok {
		return false
	}
	return true
}

func CheckUsername(username string) (b bool) {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", username); !ok {
		return false
	}
	return true
}

// GetUUID 生成唯一标识符
func GetUUID() string {
	u2 := uuid.NewV4()
	s2 := u2.String()
	res := strings.ReplaceAll(s2, "-", "")
	return res
}
