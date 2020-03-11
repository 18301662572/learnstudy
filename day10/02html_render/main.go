package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//HTML render

func loginHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"msg": "xxxlogin",
	})
}

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"msg": "xxxindex",
	})
}

func main() {
	r := gin.Default()
	//加载模板文件
	r.LoadHTMLGlob("templates/*")
	//设置静态文件的目录
	//Static()第一个参数是前台html代码里使用的路径，第二个参数是实际保存静态文件的路径
	// r.Static("/static", "./statics")
	r.Static("/dsb", "./statics")
	r.GET("/login", loginHandler)
	r.GET("/index", indexHandler)
	r.Run()
}
