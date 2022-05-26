package service

import (
	"douyin/conf"
	"fmt"
	"github.com/go-playground/assert/v2"
	"testing"
)

//func TestMain(m *testing.M) {
//	err := dao.InitDB()
//	if err != nil {
//		fmt.Printf("err:%#v\n", err)
//	} else {
//		println("connect success")
//	}
//	os.Exit(m.Run())
//}
//
//func TestRegister(t *testing.T) {
//	err := dao.InitDB()
//	if err != nil {
//		fmt.Printf("err:%#v\n", err)
//	} else {
//		println("connect success")
//	}
//	register := Register("mytest1", "123456")
//	assert.NotEqual(t, nil, register)
//	for key, value := range register {
//		fmt.Println(key, "----->", value)
//	}
//}

func TestLogin(t *testing.T) {
	conf.Init()
	conf.Redis("127.0.0.1:6379", "")
	login := Login("mytest1", "123456")
	assert.NotEqual(t, nil, login)
	for key, value := range login {
		fmt.Println(key, "----->", value)
	}
}

//func TestLogout(t *testing.T) {
//	err := dao.InitDB()
//	if err != nil {
//		fmt.Printf("err:%#v\n", err)
//	} else {
//		println("connect success")
//	}
//	logout := Logout("4186cb957db740509f1849d4ed201c90")
//	println(logout)
//}

func TestUserInfo(t *testing.T) {

	conf.Init()
	conf.Redis("127.0.0.1:6379", "")
	//login := Login("mytest1", "123456")
	//c := conf.RedisConn.Pool.Get()
	//var key = "ticket:" + login["ticket"]
	//err := c.Send("GET", key)
	//if err != nil {
	//	fmt.Printf("%s\n", err.Error())
	//}
	//err = c.Flush()
	//if err != nil {
	//	fmt.Printf("%s\n", err.Error())
	//}
	//var vb []byte
	//vb, err = redis.Bytes(c.Receive())
	//if err != nil {
	//	fmt.Printf("%s\n", err.Error())
	//}
	//fmt.Printf("%#v\n", vb)
	//l := new(model.LoginTicket)
	//// 直接将　user set,再Unmarshal为结构体时会失败
	//err = json.Unmarshal(vb, &l)
	//if err != nil {
	//	fmt.Printf("%s\n", err.Error())
	//}
	//fmt.Printf("%#v\n", l)
	info, _ := UserInfo("4e38100d07f74e178cc5ce93a73924b2")
	assert.NotEqual(t, nil, info)
	for key, value := range info {
		fmt.Println(key, "----->", value)
	}
}

//if err != nil {
//	fmt.Printf("err:%#v\n", err)
//} else {
//	println("connect success")
//}
//info, _ := UserInfo("4186cb957db740509f1849d4ed201c90")
//assert.NotEqual(t, nil, info)
//for key, value := range info {
//	fmt.Println(key, "----->", value)
//}
