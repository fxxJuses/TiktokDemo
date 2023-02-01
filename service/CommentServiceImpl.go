package service

import (
	"douyin/dao"
	"douyin/models"
)

type CommentServiceImpl struct {
}

func (csi *CommentServiceImpl) SaveComment(comment *models.Comment) error {
	err := dao.SaveComment(comment)
	return err
}

func (csi *CommentServiceImpl) FindCommentByCommentId(commentId int64) (models.Comment, error) {
	comment, err := dao.FindCommentByCommentId(commentId)

	return comment, err
}

func (csi *CommentServiceImpl) DeletComment(comment *models.Comment) error {
	err := dao.DeletComment(comment)

	return err
}

func (csi *CommentServiceImpl) FindAllCommentByVideoId(videoId int64) ([]models.Comment, error) {
	var comments []models.Comment
	err := dao.FindAllCommentByVideoId(&comments, videoId)

	return comments, err
}
