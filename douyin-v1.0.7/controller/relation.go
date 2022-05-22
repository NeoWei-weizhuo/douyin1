package controller

import (
	"douyin/model"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserListResponse struct {
	model.Response
	UserList []model.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	c.JSON(http.StatusOK, service.RelationAction(c, token))
}

// FollowList all users have same followed list
func FollowList(c *gin.Context) {
	c.JSON(http.StatusOK, service.GetUserList(c, 1))
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, service.GetUserList(c, 2))
}

// RelationAction no practical effect, just check if token is valid
//func RelationAction(c *gin.Context) {
//	token := c.Query("token")
//
//	if _, exist := usersLoginInfo[token]; exist {
//		c.JSON(http.StatusOK, model.Response{StatusCode: 0})
//	} else {
//		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
//	}
//}
//
//// FollowList all users have same follow list
//func FollowList(c *gin.Context) {
//	c.JSON(http.StatusOK, UserListResponse{
//		Response: model.Response{
//			StatusCode: 0,
//		},
//		UserList: []model.User_t{DemoUser},
//	})
//}
//
//// FollowerList all users have same follower list
//func FollowerList(c *gin.Context) {
//	c.JSON(http.StatusOK, UserListResponse{
//		Response: model.Response{
//			StatusCode: 0,
//		},
//		UserList: []model.User_t{DemoUser},
//	})
//}
