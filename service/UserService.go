package service

import (
	"douyin/models"
)

type UserService interface {
	/*

		根据用户名字查询用户信息

	*/
	GetTableUserByUsername(name string) models.User

	CreateTableUser(name string, password string) error
}
