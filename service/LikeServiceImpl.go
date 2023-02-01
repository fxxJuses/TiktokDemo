package service

import (
	"douyin/dao"
	"douyin/models"
)

type LikeServiceImpl struct {
}

func (lsi *LikeServiceImpl) GetLikeByVideoIdAndUserId(videoId string, userId int64) models.Like {
	like := dao.GetLikeByVideoIdAndUserId(videoId, userId)
	return like
}

func (lsi *LikeServiceImpl) CheckLikeByVideoIdAndUserId(videoId string, userId int64) bool {
	check := dao.CheckLikeByVideoIdAndUserId(videoId, userId)
	return check
}

func (lsi *LikeServiceImpl) UpdateLike(like *models.Like) error {
	check := dao.UpdateLike(like)
	return check
}

func (lsi *LikeServiceImpl) CreateLike(like *models.Like) error {
	check := dao.CreateLike(like)
	return check
}

func (lsi *LikeServiceImpl) FindLike(like *models.Like) bool {
	exist := dao.FindLike(like)
	return exist
}

func (lsi *LikeServiceImpl) GetLikeMapByUserId(userId int64) []models.Like {
	likes := dao.GetLikeMapByUserId(userId)
	return likes
}
