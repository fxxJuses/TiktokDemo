package service

import (
	"douyin/dao"
	"douyin/models"
)

type RelationServiceImpl struct {
}

func (rsi *RelationServiceImpl) GetRelation(relation *models.Relation) bool {
	exist := dao.GetRelation(relation)
	return exist

}

func (rsi *RelationServiceImpl) UpdateRelation(relation *models.Relation) error {
	err := dao.UpdateRelation(relation)
	return err
}

func (rsi *RelationServiceImpl) CreateRelation(relation *models.Relation) error {
	err := dao.CreateRelation(relation)
	return err
}

func (rsi *RelationServiceImpl) GetFollwByUserId(userId string) ([]models.User, error) {
	follows, err := dao.GetFollwByUserId(userId)
	if err != nil {
		return nil, err
	}
	if len(follows) == 0 {

		return nil, nil
	}
	follwIds := make([]int64, len(follows))
	for i := 0; i < len(follows); i++ {
		follwIds[i] = follows[i].ToUserId
	}

	followUsers := dao.GetFollowUsers(follwIds)
	for i := 0; i < len(followUsers); i++ {
		followUsers[i].IsFollow = true
	}

	return followUsers, err
}

func (rsi *RelationServiceImpl) GetFollwerByUserId(userId string) ([]models.User, error) {
	follows, err := dao.GetFollwerByUserId(userId)
	if err != nil {
		return nil, err
	}
	if len(follows) == 0 {

		return nil, nil
	}
	follwIds := make([]int64, len(follows))
	for i := 0; i < len(follows); i++ {
		follwIds[i] = follows[i].UserId
	}

	followUsers := dao.GetFollowUsers(follwIds)

	return followUsers, err
}
