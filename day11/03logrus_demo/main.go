package main

import (
	"github.com/sirupsen/logrus"
)

//logrus 示例

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{}) //设置JSON格式
	logrus.SetReportCaller(true)
	logrus.WithFields(logrus.Fields{
		"name": "关大门",
		"age":  80000,
	}).Warn("这是一条Warn级别的日志")
	logrus.Info("这是一条普通的日志")
}
