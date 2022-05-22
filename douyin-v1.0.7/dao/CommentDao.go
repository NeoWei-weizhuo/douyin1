package dao

import (
	"douyin/conf"
	"douyin/model"
	"gorm.io/gorm"
	"math"
	"time"
)

func DoComment(userId, videoId int64, commentText string, adder float32) bool {
	if math.Abs(float64(adder)) != 1 {
		return false
	}
	db := conf.DB
	var video *model.Video
	var comment *model.CommentVideoT
	err := db.Model(&video).Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count + ?", adder)).Error
	if err != nil {
		return false
	}

	if adder == 1 {
		err = db.Create(&model.CommentVideoT{
			UserId:      userId,
			VideoId:     videoId,
			CommentText: commentText,
			CommentTime: time.Now(),
		}).Error
	} else {
		err = db.Where("user_id = ?", userId).Where("video_id = ?", videoId).Delete(&comment).Error
	}
	return err == nil
}

func GetCommentsList(videoId int64) []model.Comment {
	db := conf.DB
	var commentList []model.CommentVideoT
	var comments []model.Comment

	db.Where("video_id = ?", videoId).Find(&commentList)

	//convert
	for _, v := range commentList {
		var comment model.Comment
		comment.Id = v.Id
		comment.Content = v.CommentText
		comment.CreateDate = v.CommentTime.String()
		//to be optimized
		err := db.Where("user_id = ?", v.UserId).Find(&comment.User).Error
		if err != nil {
		}
		comments = append(comments, comment)
	}

	return comments
}
