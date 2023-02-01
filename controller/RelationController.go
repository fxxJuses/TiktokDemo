package controller

import (
	"douyin/models"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func RelationAction(c *gin.Context) {
	// TODO 用户鉴权
	token := c.Query("token")
	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	action, _ := strconv.Atoi(c.Query("action_type"))

	// 获取当前用户 ID
	usi := service.UserServiceImpl{}
	_, user := usi.GetTableUserByToken(token)
	userId := user.Id

	relation := models.Relation{
		ToUserId: toUserId,
		UserId:   userId,
		Cancel:   action,
	}

	rsi := service.RelationServiceImpl{}
	exist := rsi.GetRelation(&relation)
	if exist {
		// 说明 关系早已存在
		// 那么就要进行更新
		relation.Cancel = action
		log.Println(relation)
		log.Println(action)
		err := rsi.UpdateRelation(&relation)
		if err != nil {
			c.JSON(http.StatusOK, models.Response{StatusCode: 1, StatusMsg: "关注失败，请重试"})
		} else {
			c.JSON(http.StatusOK, models.Response{StatusCode: 0})
		}
	} else {
		// 说明 关系不存在 需要创建关系
		err := rsi.CreateRelation(&relation)
		if err != nil {
			c.JSON(http.StatusOK, models.Response{StatusCode: 1, StatusMsg: "创建关注失败，请重试"})
		} else {
			c.JSON(http.StatusOK, models.Response{StatusCode: 0})
		}
	}
}

func FollowList(c *gin.Context) {
	userId := c.Query("user_id")
	// TODO 鉴权
	rsi := service.RelationServiceImpl{}
	users, err := rsi.GetFollwByUserId(userId)
	if err != nil {
		c.JSON(http.StatusOK, models.Response{StatusCode: 1, StatusMsg: "获取关注列表失败"})
		return
	}
	c.JSON(http.StatusOK, models.UserListResponse{
		Response: models.Response{
			StatusCode: 0,
		},
		UserList: users,
	})
}
func FollowerList(c *gin.Context) {
	userId := c.Query("user_id")
	// TODO 鉴权
	rsi := service.RelationServiceImpl{}
	users, err := rsi.GetFollwerByUserId(userId) // 获取粉丝列表
	if err != nil {
		c.JSON(http.StatusOK, models.Response{StatusCode: 1, StatusMsg: "获取粉丝列表失败"})
		return
	}
	follwUsers, err := rsi.GetFollwByUserId(userId)
	m := make(map[int64]int64)
	for i := 0; i < len(follwUsers); i++ {
		m[follwUsers[i].Id] = follwUsers[i].Id
	}
	for i := 0; i < len(users); i++ {
		_, ok := m[users[i].Id]
		if ok {
			users[i].IsFollow = true
		}
	}

	c.JSON(http.StatusOK, models.UserListResponse{
		Response: models.Response{
			StatusCode: 0,
		},
		UserList: users,
	})
}
