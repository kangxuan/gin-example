package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

type Todo struct {
	ID     int    `gorm:"primaryKey;not null" json:"id"`
	Title  string `gorm:"type:varchar(100);not null;default:'';comment:代办标题" json:"title"`
	Status bool   `json:"status"`
	IsDel  uint8  `gorm:"not null;default:0" json:"is_del"`
}

func main() {
	// 连接数据库
	db, err := gorm.Open(mysql.Open("root:123456@(192.168.10.33:3306)/test?charset=utf8mb4"), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败")
	}
	// 表迁移
	_ = db.AutoMigrate(&Todo{})

	// 实例化gin实例
	r := gin.Default()

	// 静态文件
	r.Static("/static", "./main21Templates/static")
	// 解析模板
	r.LoadHTMLFiles("./main21Templates/index.html")

	r.GET("/main21/index", func(c *gin.Context) { // 渲染模板
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/todo", func(c *gin.Context) {
			// 读取数据
			var todos []Todo
			if err := db.Where("is_del = ?", 0).Find(&todos).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, todos)
		})

		v1.POST("/todo", func(c *gin.Context) {
			// 接收JSON参数
			var addTodo Todo
			_ = c.BindJSON(&addTodo)

			if err := db.Create(&addTodo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
				return
			}

			// 返回成功
			c.JSON(http.StatusOK, gin.H{
				"message": "添加成功",
			})
		})

		// 修改操作
		v1.PUT("/todo/:id", func(c *gin.Context) {
			id := c.Param("id")

			var upTodo Todo
			// 先判断todo是否存在
			result := db.Where("id = ?", id).First(&upTodo)
			if result.RowsAffected != 1 {
				c.JSON(http.StatusNotFound, gin.H{
					"message": "没有找到此代办",
				})
				return
			}

			// 将参数进行绑定
			_ = c.BindJSON(&upTodo)
			if err := db.Save(&upTodo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "修改成功",
			})
		})

		// 删除操作
		v1.DELETE("/todo/:id", func(c *gin.Context) {
			id := c.Param("id")
			db.Model(&Todo{}).Where("id = ?", id).Update("is_del", 1)
			c.JSON(http.StatusOK, gin.H{
				"message": "删除成功",
			})
		})
	}

	_ = r.Run(":9090")
}
