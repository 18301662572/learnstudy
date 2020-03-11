package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//use logrus in gin

var log = logrus.New() //在该程序的所有文件里都能使用

func initLogrus() (err error) {
	//初始化logrus的配置
	log.Formatter = &logrus.JSONFormatter{} //记录JSON格式的日志
	//指定日志输出
	file, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed,err:", err.Error())
		return
	}
	log.Out = file
	//告诉Gin框架把他的日志也记录到我们打开的文件中
	gin.SetMode(gin.ReleaseMode) //上线的时候设置：gin设置为ReleaseMode
	gin.DefaultWriter = log.Out
	//设置日志级别
	log.Level = logrus.DebugLevel
	return
}

func indexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "Hello World!",
	})
}

func main() {
	err := initLogrus()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.GET("/index", indexHandler)
	r.Run()
}
