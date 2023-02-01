package controller

import (
	"douyin/models"
	"douyin/service"
	"douyin/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func Publish(c *gin.Context) {
	token := c.PostForm("token")
	title := c.PostForm("title")
	usi := service.UserServiceImpl{}

	existToken, user := usi.GetTableUserByToken(token)

	if !existToken {
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 1,
			StatusMsg:  "User doesn`t exist",
		})
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	//log.Println("data : ", data)
	//data :  &{movie.mp4 map[Content-Disposition:[form-data; name="data"; filename="movie.mp4"] Content-Length:[318465]
	//Content-Type:[multipart/form-data]] 318465 [0 0 0 28 102 116 121 112 109 112 52 50 0 0 0 0 10 ....
	fileSuffix := path.Ext(path.Base(data.Filename)) // 文件后缀
	fileName := utils.GeneratUUID()                  // 生成唯一的文件名字
	//log.Println("finalName : ", user.ID, "_", filename) // 0 _ movie.mp4
	//saveFile := filepath.Join("http://192.168.56.102:8020/videos/", finalName)

	// 将数据通过ftp上传到服务器中
	err = service.SaveUpLoaderFile(data, fileName, fileSuffix, title, user)
	//TODO 该步骤需要设计回滚
	if err != nil {
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		StatusCode: 0,
		StatusMsg:  fileName + " uploaded successfully",
	})

}

// 根据用户来获取其已投稿的video信息
func PublishList(c *gin.Context) {

	token := c.Query("token")

	usi := service.UserServiceImpl{}
	_, user := usi.GetTableUserByToken(token)

	vsi := service.VideoServiceImpl{}
	videos := vsi.GetAllVideosByUserId(user.Id)

	c.JSON(http.StatusOK, models.VideoListResponse{
		Response: models.Response{
			StatusCode: 0,
		},
		VideoList: videos,
	})

}
