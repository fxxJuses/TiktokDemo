package dao

import (
	"douyin/models"
)

func GetLikeByVideoIdAndUserId(videoId string, userId int64) models.Like {
	var like models.Like
	Db.Table("likes").First(&like, "video_id = ?", videoId, "user_id = ?", userId)
	return like
}

func CheckLikeByVideoIdAndUserId(videoId string, userId int64) bool {
	var like models.Like
	first := Db.Table("likes").First(&like, "video_id = ?", videoId, "user_id = ?", userId)
	if first.Error != nil {
		return false
	} else {
		return true
	}
}

func UpdateLike(like *models.Like) error {
	result := Db.Model(&like).Where("id = ?", like.Id).Update("Cancel", like.Cancel)
	// db.Model(&user).Update("name", "hello")
	return result.Error
}

func CreateLike(like *models.Like) error {
	result := Db.Table("likes").Create(&like)
	return result.Error
}

func FindLike(like *models.Like) bool {
	find := Db.Table("likes").Where("video_id = ?", like.VideoId).Where("user_id = ?", like.UserId).Find(&like)
	if find.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

func SumFavoriteCountByVideoId(videoId int64) int64 {
	result := Db.Table("likes").Where("video_id = ?", videoId).Where("cancel = ?", 1).Find(&[]models.Like{})
	return result.RowsAffected
}

func GetLikeMapByUserId(userId int64) []models.Like {
	likes := []models.Like{}
	Db.Table("likes").Where("user_id = ?", userId).Where("cancel = ?", 1).Find(&likes)
	return likes
}
func SumCommentCountByVideoId(videoId int64) int64 {
	result := Db.Table("comments").Where("video_id = ?", videoId).Where(&models.Comment{Delete: 0}).Find(&[]models.Comment{})
	return result.RowsAffected
}
