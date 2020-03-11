package mylogger

import (
	"path"
	"runtime"
)

//存放一些公用的工具函数
func getCallerInfo(skip int) (fileName string, line int, funcName string) {
	//fileName 文件的全路径 pc:存放函数对象
	pc, fileName, line, ok := runtime.Caller(skip)
	if !ok {
		return
	}
	//从fileName (文件全路径)中剥离出文件名
	fileName = path.Base(fileName) //x/y/xx.txt-->xx.txt
	//根据pc拿到函数路径
	funcName = runtime.FuncForPC(pc).Name()
	funcName = path.Base(funcName)
	return
}
