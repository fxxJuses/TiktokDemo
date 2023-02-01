package dao

import (
	"douyin/models"
	"errors"
)

func SaveComment(comment *models.Comment) error {
	err := Db.Table("comments").Save(&comment).Error
	return err
}

func FindCommentByCommentId(commentId int64) (models.Comment, error) {
	var comment models.Comment
	result := Db.Table("comments").Preload("User").First(&comment, "id = ?", commentId)
	if result.RowsAffected > 0 {
		return comment, nil
	} else {
		return comment, errors.New("评论不存在")
	}
}

func DeletComment(comment *models.Comment) error {
	err := Db.Model(&comment).Table("comments").Where("comment_id = ?", comment.Id).Update("Delete", 1).Error
	return err
}

func FindAllCommentByVideoId(comments *[]models.Comment, videoId int64) error {
	err := Db.Model(&models.Comment{}).Preload("User").Where("video_id = ?", videoId).Find(&comments).Error
	return err
}
