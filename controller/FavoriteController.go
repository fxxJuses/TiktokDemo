package controller

import (
	"douyin/models"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func FavoriteAction(c *gin.Context) {
	/*
		要处理的几件事情：
			1、点击后，数据库中的 喜欢数据要发生改变，改变如下：
				如果喜欢数据存在，则修改 喜欢列表中的 喜欢操作， 喜欢为1 不喜欢为0
				如果喜欢数据不存在， 则创建喜欢数据
			2、以上操作要同步到 video 数据中，并操作 favorite_count操作
	*/

	// TODO 这个操作肯定要放到Redis 里面的，如果不放入的话 会频繁的访问数据库，导致数据库压力过大

	// TODO 鉴权操作 省略

	token := c.Query("token")
	actionString := c.Query("action_type")
	action, _ := strconv.Atoi(actionString)

	log.Println("action:", action)
	videoIdString := c.Query("video_id")

	videoId, _ := strconv.Atoi(videoIdString)

	// 获取用户ID
	usi := service.UserServiceImpl{}
	_, user := usi.GetTableUserByToken(token)
	userId := user.Id

	lsi := service.LikeServiceImpl{}
	vsi := service.VideoServiceImpl{}

	// 根据 action 的值来判断是否需要进行操作
	like := models.Like{
		UserId:  userId,
		VideoId: int64(videoId),
		Cancel:  action,
	}
	exist := lsi.FindLike(&like)
	log.Println(like)
	if !exist && action != 2 {
		// 数据不存在 并且表示喜欢
		log.Println("No like data")
		err := lsi.CreateLike(&like)
		if err != nil {
			c.JSON(http.StatusOK, models.Response{StatusCode: 1, StatusMsg: "Create like failed"})
			return
		}
	} else {
		like.Cancel = action
		err := lsi.UpdateLike(&like)
		if err != nil {
			c.JSON(http.StatusOK, models.Response{StatusCode: 1, StatusMsg: "Update like failed"})
			return
		}
	}

	// TODO 这个操作可以进行优化
	vsi.UpdateFavoriteByVideoId(int64(videoId))

	c.JSON(http.StatusOK, models.Response{StatusCode: 0})
}

func FavoriteList(c *gin.Context) {
	// TODO  鉴权
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	vsi := service.VideoServiceImpl{}

	videos := vsi.GetAllVideos()
	lsi := service.LikeServiceImpl{}
	likes := lsi.GetLikeMapByUserId(userId)

	var set map[int64]struct{}
	set = make(map[int64]struct{})

	for _, value := range likes {
		set[value.VideoId] = struct{}{}
	}
	favoriteVideos := make([]models.Video, len(likes))
	index := 0
	for i := 0; i < len(videos); i++ {
		if _, ok := set[videos[i].Id]; ok {
			favoriteVideos[index] = videos[i]
			index++
		}
	}
	log.Println(favoriteVideos)
	c.JSON(http.StatusOK, models.VideoListResponse{
		Response: models.Response{
			StatusCode: 0,
		},
		VideoList: favoriteVideos,
	})

}
