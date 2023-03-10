package controller

import (
	"douyin/models"
	"douyin/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"
)

var tempChat = map[string][]models.Message{}

var messageIdSequence = int64(1)

type ChatResponse struct {
	models.Response
	MessageList []models.Message `json:"message_list"`
}

// MessageAction no practical effect, just check if token is valid
func MessageAction(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	content := c.Query("content")
	usi := service.UserServiceImpl{}
	_, user := usi.GetTableUserByToken(token)

	userIdB, _ := strconv.Atoi(toUserId)
	chatKey := genChatKey(int64(user.Id), int64(userIdB))

	atomic.AddInt64(&messageIdSequence, 1)
	curMessage := models.Message{
		Id:         messageIdSequence,
		Content:    content,
		CreateTime: time.Now().Format(time.Kitchen),
	}

	if messages, exist := tempChat[chatKey]; exist {
		tempChat[chatKey] = append(messages, curMessage)
	} else {
		tempChat[chatKey] = []models.Message{curMessage}
	}
	c.JSON(http.StatusOK, models.Response{StatusCode: 0})

}

// MessageChat all users have same follow list
func MessageChat(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")

	usi := service.UserServiceImpl{}
	_, user := usi.GetTableUserByToken(token)

	userIdB, _ := strconv.Atoi(toUserId)
	chatKey := genChatKey(int64(user.Id), int64(userIdB))

	c.JSON(http.StatusOK, ChatResponse{Response: models.Response{StatusCode: 0}, MessageList: tempChat[chatKey]})

}

func genChatKey(userIdA int64, userIdB int64) string {
	if userIdA > userIdB {
		return fmt.Sprintf("%d_%d", userIdB, userIdA)
	}
	return fmt.Sprintf("%d_%d", userIdA, userIdB)
}
