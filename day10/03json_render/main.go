package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//json render

func indexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "这是一个index页面",
	})
}

func helloHandler(c *gin.Context) {
	var msg struct {
		Name    string `json:"user"`
		Message string
		Age     int
	}
	msg.Name = "小王子"
	msg.Message = "Hello world!"
	msg.Age = 18
	c.JSON(http.StatusOK, msg)
}

func xmlHandler(c *gin.Context) {
	c.XML(http.StatusOK, gin.H{
		"msg": "xml",
	})
}

func main() {
	r := gin.Default()
	r.GET("/index", indexHandler)
	r.GET("/hello", helloHandler)
	r.GET("/xml", xmlHandler)
	r.GET("/yaml", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"msg": "ok", "status": http.StatusOK})
	})
	r.Run()
}
