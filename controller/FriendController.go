package controller

import (
	"douyin/models"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func FriendList(c *gin.Context) {
	// TODO 不知道为什么 按照官方给的文档 返回了该返回的东西，但是Friend界面没有任何响应，
	// 发现是 Friend 按钮，没有发送请求的能力
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
	log.Println(users)
	friends := make([]models.Friend, len(users))
	for i := 0; i < len(users); i++ {
		friends[i].User = users[i]
		friends[i].MsgType = 0
		friends[i].Message = "你好啊，能看到消息吗？"
	}
	log.Println(friends)
	c.JSON(http.StatusOK, models.FriendListResponse{
		Response: models.Response{
			StatusCode: 0,
		},
		FriendList: friends,
	})

}
