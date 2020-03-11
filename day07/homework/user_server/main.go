package main

import (
	"code.oldbody.com/studygolang/learnstuday/day07/homework/mylogger"
)

//测试myloger的程序

var logger mylogger.Logger

func main() {
	logger = mylogger.NewFileLogger("debug", "xx.log", "./")
	defer logger.Close()
	for {
		logger.Debug("下雨了")
		logger.Error("打雷了")
	}

}
