package model

import (
	"github.com/13808796047/go-blog/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB gorm.DB对象

var DB *gorm.DB

// ConnectDB初始化模型
func ConnectDB() *gorm.DB {
	var err error
	config := mysql.New(mysql.Config{
		DSN: "homestead:secret@tcp(127.0.0.1:33060)/goblog?charset=utf8&parseTime=True&loc=Local",
	})
	// 准备数据库连接池
	DB, err = gorm.Open(config, &gorm.Config{})
	logger.LogError(err)
	return DB
}
