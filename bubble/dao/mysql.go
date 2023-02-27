package dao

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitMysql 初始化数据库
func InitMysql() {
	err := errors.New("")
	// 连接数据库
	DB, err = gorm.Open(mysql.Open("root:123456@(192.168.10.33:3306)/test?charset=utf8mb4"), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败")
	}

	return
}
