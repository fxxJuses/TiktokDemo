package controller

import (
	"douyin/models"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Feed(c *gin.Context) {
	token := c.Query("token")
	usi := service.UserServiceImpl{}
	_, user := usi.GetTableUserByToken(token)

	// TODO 还需要返回一个 is_favorite 的信息 需要结合 likes 表来实现
	vsi := service.VideoServiceImpl{}
	videos := vsi.GetAllVideos()
	//log.Println(videos)

	lsi := service.LikeServiceImpl{}
	likes := lsi.GetLikeMapByUserId(user.Id)

	var set map[int64]struct{}
	set = make(map[int64]struct{})

	for _, value := range likes {
		set[value.VideoId] = struct{}{}
	}

	for i := 0; i < len(videos); i++ {
		if _, ok := set[videos[i].Id]; ok {
			videos[i].IsFavorite = true
		}
	}

	c.JSON(http.StatusOK, models.FeedResponse{
		Response:  models.Response{StatusCode: 0},
		VideoList: videos,
		NextTime:  time.Now().Unix(),
	})
}
