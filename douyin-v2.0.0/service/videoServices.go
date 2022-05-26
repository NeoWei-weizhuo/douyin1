package service

import (
	"douyin/dao"
	"douyin/model"
	"time"
)

func SaveVideoToDatabase(saveFile string, userId int64, coverUrl string, title string) error {

	// 准备向数据库视频表单插入的数据
	v := model.DatabaseVideo{
		AuthorId:      userId,
		PlayUrl:       saveFile,
		CoverUrl:      coverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         title,
		Timestamp:     time.Now(),
	}
	// save video message into database
	err := dao.SaveVideoMessage(v)

	if err != nil {
		return err
	}

	return nil
}

func GetVideosByUserId(userId int64) ([]model.Video, error) {
	// 通过userId获取其所投稿的视频列表
	databaseVideos, err := dao.GetVideosByUserId(userId)
	if err != nil {
		return nil, err
	}

	// 从DatabaseVideo转换为Video，添加部分信息
	videos := make([]model.Video, 0, len(databaseVideos))
	if databaseVideos != nil && len(databaseVideos) != 0 {
		// 通过userId得到Author基本信息
		databaseUser, err := dao.GetAuthorById(userId)
		if err != nil {
			return nil, err
		}
		// 获得作者信息，作者不能自己关注自己
		author := model.User{
			Id:            databaseUser.Id,
			Username:      databaseUser.Name,
			Followcount:   databaseUser.FollowCount,
			Followercount: databaseUser.FollowerCount,
			IsFollow:      false}

		for _, v := range databaseVideos {
			// 判断本人是否对该视频点赞
			isFavor, err := dao.IsFavor(v.Id, v.AuthorId)
			if err != nil {
				return nil, err
			}
			// 包装Video结构
			videos = append(videos, model.Video{Id: v.Id,
				Author: author, PlayUrl: v.PlayUrl, CoverUrl: v.CoverUrl,
				FavoriteCount: v.FavoriteCount, CommentCount: v.CommentCount,
				IsFavorite: isFavor})
		}
	}
	return videos, nil
}
