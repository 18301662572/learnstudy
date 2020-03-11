package mylog

import (
	"fmt"
	"path"
	"runtime"
)

func getCallerInfo() (filename, funcNamme string, line int) {
	pc, filename, line, ok := runtime.Caller(2) //参数从0开始，最底层为0，调用一次+1
	if !ok {
		return
	}
	//根据pc拿到当前执行的函数名
	funcNamme = runtime.FuncForPC(pc).Name()
	funcNamme = path.Base(funcNamme) //获取路径的最后一个 / 后的内容
	filename = path.Base(filename)
	fmt.Println(funcNamme, filename, line)
	return
}
