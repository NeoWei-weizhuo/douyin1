package service

import (
	"strconv"

	"douyin/dao"
	"douyin/model"
	"github.com/gin-gonic/gin"
)

func GetUserIdFromUserInfo(token string) (int64, model.Response) {
	m, _ := UserInfo(token)
	if err, exist := m["errMsg"]; exist {
		return -1, model.Response{StatusCode: -1,StatusMsg: err}

	}
	userIdString, Exists := m["userId"]
	if !Exists {
		return -1,model.Response{StatusCode: -1, StatusMsg: "用户不存在或者未登录"}
	}
	//根据登录状态获取user_id
	userId, err := strconv.ParseInt(userIdString, 10, 64)
	if err != nil {
		return -1,model.Response{StatusCode: -1, StatusMsg: "用户id信息有误"}
	}
	return userId,model.Response{StatusCode: 1, StatusMsg: "查找成功"}
}
func RelationAction(c *gin.Context, token string) model.Response {

	//用户鉴权


	// 检查token是否存在合法


	userId, resp := GetUserIdFromUserInfo(token)
	if resp.StatusCode == -1 {
		return model.Response{StatusCode: -1, StatusMsg: "操作失败"}
	}
	//关注
	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"),10,64)

	actionType, err := strconv.ParseInt(c.Query("action_type"), 10, 32)
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = "请求参数有误"
		return resp
	}

	if actionType == 1 { //关注
		if dao.DoFollow(userId, toUserId) {
			resp.StatusMsg = "关注成功"
		} else {
			resp.StatusMsg = "关注失败"
			resp.StatusCode = -1
			return resp
		}
		resp.StatusMsg = "关注成功"
	} else if actionType == 2 { //取消关注
		if dao.CancelFollow(userId, toUserId) {
			resp.StatusMsg = "取关成功"
		} else {
			resp.StatusMsg = "取关失败"
			resp.StatusCode = -1
			return resp
		}
	}

	resp.StatusCode = 0
	return resp
	//关注逻辑
	//userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	//toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	//actionType, err := strconv.ParseInt(c.Query("action_type"), 10, 32)
	//if err != nil {
	//	resp.StatusCode = -1
	//	resp.StatusMsg = "请求参数有误"
	//	return resp
	//}
	//if actionType == 1 { //关注
	//	if dao.DoFollow(userId, toUserId) {
	//		resp.StatusMsg = "关注成功"
	//	} else {
	//		resp.StatusMsg = "关注失败"
	//		resp.StatusCode = -1
	//		return resp
	//	}
	//	resp.StatusMsg = "关注成功"
	//} else if actionType == 2 { //取消关注
	//	if dao.CancelFollow(userId, toUserId) {
	//		resp.StatusMsg = "取关成功"
	//	} else {
	//		resp.StatusMsg = "取关失败"
	//		resp.StatusCode = -1
	//		return resp
	//	}
	//}
	//
	//resp.StatusCode = 0
	//return resp
}

func GetUserList(c *gin.Context, flag int) model.UserListResponse {
	var userList []model.User
	var resp model.Response

	//if flag == 1 { //1表示获取关注列表
	//	s = "user_id"
	//} else if flag == 2 { //2表示获取粉丝列表
	//	s = "to_user_id"
	//}
	//用户鉴权
	token := c.Query("token")
	userId, resp := GetUserIdFromUserInfo(token)
	if resp.StatusCode == -1 {
		return model.UserListResponse{
			Response: resp,
			UserList: []model.User{},
		}
	}



	userList = dao.GetUserListById(userId, flag)
	resp.StatusCode = 0
	return model.UserListResponse{
		Response: resp,
		UserList: userList,
	}
}
