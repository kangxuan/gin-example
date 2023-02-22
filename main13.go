package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	// 解析模板
	r.LoadHTMLFiles("./main13Templates/upload.tmpl")
	r.GET("/main13/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.tmpl", nil)
	})

	r.POST("/main13/upload", func(c *gin.Context) {
		f, err := c.FormFile("file1")
		if err != nil {
			fmt.Printf("接收文件错误：%s\n", err.Error())
			return
		}

		// 打印文件名
		log.Println(f.Filename)
		dst := fmt.Sprintf("./main13Templates/%s", f.Filename)
		// 或者通过path
		//dst := path.Join("./main13Templates/", f.Filename)
		// 上传文件
		err = c.SaveUploadedFile(f, dst)
		if err != nil {
			fmt.Printf("保存上传文件失败：%v\n", err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "uploaded！",
		})
	})

	_ = r.Run(":9090")
}
