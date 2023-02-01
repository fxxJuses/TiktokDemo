package main

import (
	"douyin/dao"
	"douyin/service"
	"github.com/gin-gonic/gin"
)

func main() {
	initDeps()

	//go service.RunMessageServer()

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
}
