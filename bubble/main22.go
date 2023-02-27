package main

import (
	dao2 "gin-example/bubble/dao"
	models2 "gin-example/bubble/models"
	routers2 "gin-example/bubble/routers"
)

func main() {
	// 初始化数据库
	dao2.InitMysql()
	// 表迁移
	_ = dao2.DB.AutoMigrate(&models2.Todo{})
	// 注册路由
	r := routers2.SetupRouter()
	// 运行项目
	_ = r.Run(":9090")
}
