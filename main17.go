package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

// Gorm 基础
func main() {
	// 连接数据库
	db, err := gorm.Open(mysql.Open("root:123456@(192.168.10.33:3306)/test?charset=utf8mb4"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database!")
	}

	// 自动生成表 迁移 schema
	db.AutoMigrate(&Product{})

	// 写入新数据
	//db.Create(&Product{
	//	Code:  "C1122334455",
	//	Price: 190000,
	//})

	// 查询数据
	var p1 Product
	//db.First(&p1, 1)
	//db.First(&p1, "code = ?", "C1122334455") // 按查询条件查询
	//fmt.Printf("record: %d\n", p1.ID)

	// 更新数据
	//db.Model(&p1).Update("Price", 4000)                         // 更新单个字段
	//db.Model(&p1).Updates(Product{Code: "DS111", Price: 45000}) // 更新多个字段
	//db.Model(&p1).Updates(map[string]interface{}{
	//	"Code":  "DS112",
	//	"Price": 2409090,
	//})

	// 删除数据
	db.Delete(&p1, 2)
}
