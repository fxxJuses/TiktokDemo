package test

import (
	"douyin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"testing"
)

func TestCreateTable(t *testing.T) {
	var err error
	var db *gorm.DB
	dsn := "root:0102kyrie@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	//想要正确的处理time.Time,需要带上 parseTime 参数，
	//要支持完整的UTF-8编码，需要将 charset=utf8 更改为 charset=utf8mb4
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln("err:", err.Error())
	}

	// 迁移 schema
	db.AutoMigrate(&models.User{})

	t.Log("hello world")

}
