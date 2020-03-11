package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func uploadHandler(c *gin.Context) {
	//提取用户上传的文件
	fileobj, err := c.FormFile("filename")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}
	//fileobj 上传的文件对象
	//fileobj.Filename 上传文件名
	filePath := fmt.Sprintf("./%s", fileobj.Filename)
	//保存文件到本地路径
	c.SaveUploadedFile(fileobj, filePath)
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

//上传文件示例
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./upload.html")
	r.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", nil)
	})

	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	r.POST("/upload", uploadHandler)
	r.Run()
}
