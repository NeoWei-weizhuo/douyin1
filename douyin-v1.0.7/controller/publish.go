package controller

import (
	"bytes"
	"douyin/conf"
	"douyin/model"
	"douyin/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"image"
	"image/jpeg"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

type VideoListResponse struct {
	model.Response
	VideoList []model.Video `json:"video_list"`
}

// getCover videoPath为视频的存储路径，fileName为userId和时间戳拼接的封面名
func getCover(videoPath string, fileName string) (string, error) {
	// 封面的宽高
	width := 460
	height := 600
	// 封面的存储路径
	coverName := filepath.Join("./public/cover/",
		fileName+".jpg")
	// 封面在数据库中存储的url
	coverurl := "http://" + conf.HostAndPort() + "/static/cover/" + fileName + ".jpg"

	cmd := exec.Command("ffmpeg", "-i", videoPath,
		"-vframes", "1", "-s", fmt.Sprintf("%dx%d", width, height),
		"-f", "singlejpeg", "-")
	var buffer bytes.Buffer
	cmd.Stdout = &buffer
	if err := cmd.Run(); err != nil {
		fmt.Println("获取视频第一帧失败")
		return "", err
	}

	// 在获得buffer后存入jpg图片
	m, _, _ := image.Decode(&buffer)
	f, _ := os.Create(coverName)
	defer f.Close()
	err := jpeg.Encode(f, m, nil)
	if err != nil {
		fmt.Println("生成封面失败")
		return "", err
	}

	return coverurl, nil
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	// 获得token和title
	token := c.PostForm("token")
	title := c.PostForm("title")

	// 检查token是否存在合法
	m, _ := service.UserInfo(token)
	if err, exist := m["errMsg"]; exist {
		c.JSON(http.StatusOK,
			model.Response{StatusCode: 1,
				StatusMsg: err})
		return
	}
	userIdString, Exists := m["userId"]
	if !Exists {
		c.JSON(http.StatusOK,
			model.Response{StatusCode: 1,
				StatusMsg: "用户不存在或未登录"})
		return
	}
	userId, err := strconv.ParseInt(userIdString, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK,
			model.Response{StatusCode: 1,
				StatusMsg: "userId不为整数"})
		return
	}

	// 获得上传的视频
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 生成视频存放路径，并且保存视频
	filename := filepath.Base(data.Filename)
	// 获得当前时间戳，避免用户重复上传同一视频导致的混乱
	timestamp := time.Now().Unix()
	finalName := fmt.Sprintf("%d_%d_%s", userId, timestamp, filename)
	// 视频保存到./public文件夹下
	saveFile := filepath.Join("./public/", finalName)
	// 生成视频存储的url
	databaseFileUrl := "http://" + conf.HostAndPort() + "/static/" + finalName
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 通过获得的视频生成封面信息，截取视频第一帧作为封面
	coverUrl, err := getCover(saveFile, fmt.Sprintf("%d_%d_cover", userId, timestamp))
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 将视频信息写入数据库
	err = service.SaveVideoToDatabase(databaseFileUrl, userId, coverUrl, title)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		StatusCode: 0,
		StatusMsg:  finalName + "上传成功",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	// 获得token和user_id
	token := c.Query("token")
	user_id_string := c.Query("user_id")

	// 检查token的合法性
	m, _ := service.UserInfo(token)
	if err, exist := m["errMsg"]; exist {
		c.JSON(http.StatusOK,
			VideoListResponse{
				Response: model.Response{
					StatusCode: 1,
					StatusMsg:  err,
				},
			})
		return
	}
	// 比较token和user_id的一致性
	userIdString, Exists := m["userId"]
	if !Exists {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: model.Response{
				StatusCode: 1,
				StatusMsg:  "user not exists",
			},
			VideoList: []model.Video{},
		})
		return
	}
	userId, _ := strconv.ParseInt(userIdString, 10, 64)
	user_id, err := strconv.ParseInt(user_id_string, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: model.Response{
				StatusCode: 1,
				StatusMsg:  "user_id不是整数",
			},
		})
		return
	}
	if userId != user_id {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: model.Response{
				StatusCode: 1,
				StatusMsg:  "token与user_id不匹配",
			},
		})
		return
	}

	// 通过userId获得该用户所有视频的列表
	videos, err := service.GetVideosByUserId(userId)
	if err != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: model.Response{
				StatusCode: 1,
				StatusMsg:  "无法得到用户视频列表",
			},
		})
		return
	}

	c.JSON(http.StatusOK, VideoListResponse{
		Response: model.Response{
			StatusCode: 0,
		},
		VideoList: videos,
	})
}
