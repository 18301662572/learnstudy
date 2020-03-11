package main

import (
	"fmt"

	"code.oldbody.com/studygolang/learnstuday/day04/14log/mylog"
)

func main() {
	f1 := mylog.NewFileLogger(mylog.DEBUG, "./", "test.log")
	f1.Debug("这是一个测试日志")
	fmt.Println("可以申请IPO")
	f1.Debug("这是一条Info级别的日志")
	fmt.Println()

	userID := 10
	f1.Debug("id是%d的用户一直在尝试登陆\n", userID)
}
