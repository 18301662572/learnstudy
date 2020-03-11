package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//参数相关示例

func queryStringHandler(c *gin.Context) {
	//获取query string参数
	// nameValue := c.Query("name") //查不到默认是空字符串
	nameValue := c.DefaultQuery("name", "找不到名字") //查不到就用指定的默认值（第二个参数）
	cityValue := c.Query("city")
	c.JSON(http.StatusOK, gin.H{
		"name": nameValue,
		"city": cityValue,
	})
}

func formHandler(c *gin.Context) {
	//提取form表单的数据
	nameValue := c.PostForm("name")
	cityValue := c.DefaultPostForm("city", "未知城市")
	c.JSON(http.StatusOK, gin.H{
		"name": nameValue,
		"city": cityValue,
	})
}

func paramsHandler(c *gin.Context) {
	//提取路径参数
	actionValue := c.Param("action")
	c.JSON(http.StatusOK, gin.H{
		"action": actionValue,
	})
}

func main() {
	r := gin.Default()
	//query string: http://127.0.0.1:8080/search?name=小王子&city=beijing
	r.GET("/query_string", queryStringHandler)
	//form params:html页面上form表单的提交数据
	r.POST("/form", formHandler)

	//获取 path参数/URL参数： /book/list  /book/new   /book/delete
	//                      /posts/2019/01   /posts/2019/06
	r.GET("/book/:action", paramsHandler)
	r.Run()
}
