package mylog

import (
	"fmt"
	"os"
	"time"
)

//FileLogger 往文件中记录日志的结构体
type FileLogger struct {
	level       int //只有大于这个级别的日志才记录，小于的就不管了
	logFilePath string
	logFileName string
	logFile     *os.File //os包中File类型的指针
}

//NewFileLogger 是一个生成文件日志结构体实例的构造函数
func NewFileLogger(level int, logFilePath, logFileName string) *FileLogger {
	flobj := &FileLogger{
		level:       level,
		logFileName: logFileName,
		logFilePath: logFilePath,
	}
	flobj.initFileLogger() //调用下面的初始化文件句柄的方法
	return flobj
}

//用来初始化文件日志的文件句柄
func (f *FileLogger) initFileLogger() {
	filepath := fmt.Sprintf("%s/%s", f.logFilePath, f.logFileName)
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Sprintf("open file:%s failed", filepath))
	}
	f.logFile = file //把日志文件复制给结构体中loglogFile这个字段
}

//记录日志
func (f *FileLogger) Debug(format string, args ...interface{}) {
	if f.level > DEBUG { //如果设置的日志级别大于当前级别，不用写日志
		return
	}
	filename, funcNamme, line := getCallerInfo()
	//往文件里面写
	//f.logFile.WriteString(msg)
	//日志的格式 时间 日志级别 哪个文件哪行哪个函数 日志信息
	//[2019-7-19 19:20:59] [DEBUG] main.go [14] msg
	nowStr := time.Now().Format("[2006-01-02 15:04:05.000]")
	format = fmt.Sprintf("%s [%s] [%s:%s] [%d] %s",
		nowStr, getlevelStr(f.level), filename, funcNamme, line, format)
	fmt.Fprintf(f.logFile, format, args...)
	fmt.Fprintln(f.logFile) //加换行
}

func (f *FileLogger) Info(msg string) {
	//往文件里面写
	//往哪个文件里面写？
	f.logFile.WriteString(msg)

}

func (f *FileLogger) Error(msg string) {
	//往文件里面写
	//往哪个文件里面写？
	f.logFile.WriteString(msg)
}
