package controller

import (
	"douyin/models"
	"douyin/service"
	"douyin/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := utils.InputPassDBPass(username+password, "1a2b3c4d")

	usi := service.UserServiceImpl{}

	user, err := usi.CreateTableUser(username, token)

	if err != nil {
		c.JSON(http.StatusOK, models.UserLoginResponse{
			Response: models.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		c.JSON(http.StatusOK, models.UserLoginResponse{
			Response: models.Response{StatusCode: 0},
			UserId:   int64(user.Id),
			Token:    user.Password,
		})
	}

}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := utils.InputPassDBPass(username+password, "1a2b3c4d")

	var usi = service.UserServiceImpl{}
	user, err := usi.GetTableUserByUsername(username)
	log.Println(err)
	if err != nil {
		c.JSON(http.StatusOK, models.UserLoginResponse{
			Response: models.Response{StatusCode: 1, StatusMsg: "User don`t exist"},
		})
	} else if strings.Compare(user.Password, token) == 0 {
		c.JSON(http.StatusOK, models.UserLoginResponse{
			Response: models.Response{StatusCode: 0},
			UserId:   int64(user.Id),
			Token:    token,
		})
	} else {

		c.JSON(http.StatusOK, models.UserLoginResponse{
			Response: models.Response{StatusCode: 1, StatusMsg: "password is not correct"},
		})
	}
}

func UserInfo(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))

	token := c.Query("token")

	usi := service.UserServiceImpl{}
	byToken, user := usi.GetTableUserByToken(token)
	if (byToken) && (userId == int(user.Id)) {
		c.JSON(http.StatusOK, models.UserResponse{
			Response: models.Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, models.UserResponse{
			Response: models.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}

}
