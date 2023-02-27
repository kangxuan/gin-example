package controllers

import (
	models2 "gin-example/bubble/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// IndexHandler 首页
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

// GetGotoList 获取代办列表
func GetGotoList(c *gin.Context) {
	todoList, err := models2.GetAllGoto()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, todoList)
}

// AddGoto 添加代办
func AddGoto(c *gin.Context) {
	// 接收JSON参数
	var addTodo models2.Todo
	_ = c.BindJSON(&addTodo)

	if err := models2.AddAGoto(&addTodo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	// 返回成功
	c.JSON(http.StatusOK, gin.H{
		"message": "添加成功",
	})
}

// UpdateGoto 更新代办
func UpdateGoto(c *gin.Context) {
	id := c.Param("id")

	upTodo, err := models2.GetAGotoById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "没有找到此代办",
		})
		return
	}
	// 将参数进行绑定
	_ = c.BindJSON(upTodo)

	if err = models2.UpdateAGoto(upTodo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "修改成功",
	})
}

// DeleteGoto 删除代办
func DeleteGoto(c *gin.Context) {
	id := c.Param("id")

	if err := models2.DeleteAGoto(id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "删除失败",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
}
