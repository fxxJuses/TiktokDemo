package dao

import (
	"douyin/models"
)

func GetRelation(relation *models.Relation) bool {
	result := Db.Table("relations").First(&relation)
	if result.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

func UpdateRelation(relation *models.Relation) error {
	err := Db.Model(&models.Relation{}).Where("id = ?", relation.Id).Update("cancel", relation.Cancel).Error
	return err
}

func CreateRelation(relation *models.Relation) error {
	err := Db.Table("relations").Create(&relation).Error
	return err
}

func GetFollwByUserId(userId string) ([]models.Relation, error) {
	var relations []models.Relation
	result := Db.Table("relations").Where("user_id = ? AND cancel = ?", userId, 1).Find(&relations)
	return relations, result.Error
}

func GetFollwerByUserId(userId string) ([]models.Relation, error) {
	var relations []models.Relation
	result := Db.Table("relations").Where("to_user_id = ? AND cancel = ?", userId, 1).Find(&relations)
	return relations, result.Error
}
