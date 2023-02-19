package main

import (
	"douyin/dao"
	"douyin/middleware/rabbitMQ"
	"douyin/middleware/redis"
	"douyin/service"
	"douyin/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	initDeps()

	go service.RunMessageServer()

	r := gin.Default()

	initRouter(r)

	err := r.Run(":8020")
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func initDeps() {
	dao.InitDao()
	service.InitFtp()
	err := redis.InitClient()
	if err != nil {
		log.Println("redis init falied, please check your redis conf")
		return
	}
	rabbitMQ.InitRabbitMQ()
	// 敏感词库加载
	utils.FilterInit()
}
