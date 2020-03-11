package main

import (
	"code.oldbody.com/studygolang/learnstuday/day05/mylogger"
)

//一个使用自定义日志库的用户程序

//接口的使用
var logger mylogger.Logger

func main() {
	//1.往文件里面写日志
	logger = mylogger.NewFileLogger("info", "./", "xxx.log")
	//关闭文件句柄。实现接口，实现Close()方法
	defer logger.Close()
	//2.往终端打印日志
	//logger = mylogger.NewConsoleLogger("debug")
	for {
		sb := "管大妈"
		logger.Debug("%s是个好捧哏", sb)
		logger.Info("Info")
		logger.Warn("Warn")
	}
}

//作业
//按照时间切分（一个小时一切）
//文件日志的结构体里面记录一下上一次切分的小时
//每一次写日志的时候 拿当前的时间的小时数和上一次切分的小时数作比较
//如果小时数不一致就切分,否则就不切分
