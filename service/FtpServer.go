package service

import (
	"douyin/dao"
	"douyin/models"
	"douyin/utils"
	"errors"
	"gopkg.in/dutchcoders/goftp.v1"
	"log"
	"mime/multipart"
	"os"
	"path"
)

var ftp *goftp.FTP

func InitFtp() {
	var err error

	// For debug messages: goftp.ConnectDbg("ftp.server.com:21")  # 这个地方的连接时间太长了，需要优化
	if ftp, err = goftp.Connect("192.168.56.102:21"); err != nil {
		log.Println(err)

	}

	// Username / password authentication
	if err = ftp.Login("ftpuser", "123456"); err != nil {
		log.Println(err)

	} else {
		log.Println("Successfully connected !!")
	}

}

func SaveUpLoaderFile(data *multipart.FileHeader, fileName string, fileSuffix string, title string, user models.User) error {
	// 现将mp4 上传到ftp服务器中
	file, err := data.Open()
	err = ftp.Stor(path.Join("/videos", fileName+fileSuffix), file)
	if err != nil {
		log.Println("上传视频文件失败")
		return errors.New("上传视频文件失败")
	}
	videoPath := "http://192.168.56.102:8020/videos/" + fileName + fileSuffix
	snapShotPath := "./temp/images/" + fileName + ".jpg"
	utils.GetSnapshot(videoPath, snapShotPath, 1)
	imageFile, err := os.Open(snapShotPath)
	if err != nil {
		log.Println("图片格式错误")
		return errors.New("图片格式错误")
	}

	err = ftp.Stor(path.Join("/images", fileName+".jpg"), imageFile)
	if err != nil {
		log.Println("上传视频截图失败")
		return errors.New("上传视频截图失败")
	}

	videoInfo := models.VideoTable{
		AuthorId:      int64(user.Id),
		PlayUrl:       videoPath,
		CoverUrl:      "http://192.168.56.102:8020/images/" + fileName + ".jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         title,
	}

	err = dao.UpLoadInfoInsert(&videoInfo)

	return err
}
