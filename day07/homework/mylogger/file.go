package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

//往文件里面写日志的代码，按照大小切分

//logData 日志信息结构体
type logData struct {
	Message  string
	LogLevel string
	LineNo   int
	TimeStr  string
	FuncName string
	FileName string
}

//FileLogger 文件日志结构体
type FileLogger struct {
	level       Level //日志级别
	fileName    string
	filePath    string
	file        *os.File
	errFile     *os.File
	maxSize     int64         //记录日志大小字段
	logDataChan chan *logData //定义一个存放日志信息的通道
}

// NewFileLogger 文件日志结构体的构造函数
func NewFileLogger(levelStr, fileName, filePath string) *FileLogger {
	logLevel := parseLogLevel(levelStr)
	fl := &FileLogger{
		fileName:    fileName,
		filePath:    filePath,
		level:       logLevel,
		maxSize:     10 * 1024 * 1024,
		logDataChan: make(chan *logData, 50000), //通道初始化
	}
	fl.initFile() //根据上面的文件路径和文件名打开日志文件，把文件句柄赋值给结构体对应的字段

	return fl
}

//将指定的日志文件打开赋值给结构体
func (f *FileLogger) initFile() {
	logName := path.Join(f.filePath, f.fileName)
	//打开文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		panic(fmt.Errorf("打开日志文件%s失败,%v", logName, err))
	}
	f.file = fileObj

	//打开错误日志的文件
	errlogName := fmt.Sprintf("%s.err", logName)
	errFileObj, err := os.OpenFile(errlogName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		panic(fmt.Errorf("打开错误日志文件%s失败,%v", errlogName, err))
	}
	f.errFile = errFileObj

	//开启goroutine去写日志
	go f.writeLogBackgoud()

}

//检查是否要拆分,检查当前日志文件的大小是否超过了maxSize
func (f *FileLogger) checkSplit(file *os.File) bool {
	fileInfo, _ := file.Stat() //拿到文件的所有信息
	fileSize := fileInfo.Size()
	return fileSize >= f.maxSize //当传进来的日志文件的大小超过限额就返回true
}

//封装一个切分日志文件的方法
func (f *FileLogger) splitLogFile(file *os.File) *os.File {
	//切分文件
	fileName := file.Name() //拿到文件的完整路径
	backupName := fmt.Sprintf("%s_%v.back", fileName, time.Now().Unix())
	//1.把原来的文件关闭
	file.Close()
	//2.备份原来的文件(重命名原来的文件)
	os.Rename(fileName, backupName)
	//3.新建一个文件
	fileObj, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		panic(fmt.Errorf("打开日志文件失败"))
	}
	return fileObj
}

//将公用的记录日志的功能封装成一个单独的方法
func (f *FileLogger) log(level Level, format string, args ...interface{}) {
	if f.level > level {
		return
	}
	// f.file.Write()
	msg := fmt.Sprintf(format, args...) //得到用户要记录的日志
	//日志格式：[时间][文件：行号][函数名][日志级别]  日志信息
	nowStr := time.Now().Format("2006-01-02 15:04:05.000")
	fileName, line, funcName := getCallerInfo(3)
	logLevelStr := getLevelStr(level)

	//将日志信息发送到通道中，以备写日志的goroutine去接收值
	//1.构造日志数据结构体
	logData := &logData{
		Message:  msg,
		LogLevel: logLevelStr,
		LineNo:   line,
		TimeStr:  nowStr,
		FuncName: funcName,
		FileName: fileName,
	}
	//当通道满了，会阻塞
	select {
	case f.logDataChan <- logData:
	default:
		<-f.logDataChan          //丢弃最早的一条
		f.logDataChan <- logData //把最新的发送进去
	}

}

//真正往文件里面写日志的函数
func (f *FileLogger) writeLogBackgoud() {
	for data := range f.logDataChan {
		level := parseLogLevel(data.LogLevel)
		logMsg := fmt.Sprintf("[%s][%s:%d][%s][%s]%s",
			data.TimeStr, data.FileName, data.LineNo, data.FuncName, data.LogLevel, data.Message)
		//往文件里写之前要做一个检查
		//检查当前日志文件的大小是否超过了maxSize
		if f.checkSplit(f.file) {
			//切分文件
			f.file = f.splitLogFile(f.file)
		}
		fmt.Fprintln(f.file, logMsg) //利用fmt包将msg字符串写入f.file文件中
		//如果是error或fatal级别的日志还要记录到f.errFile中
		if level >= ErrorLevel {
			if f.checkSplit(f.errFile) {
				f.errFile = f.splitLogFile(f.errFile)
			}
			fmt.Fprintln(f.errFile, logMsg)
		}
	}
}

//Debug 方法
func (f *FileLogger) Debug(format string, args ...interface{}) {
	f.log(DebugLevel, format, args...)
}

//Warn 方法
func (f *FileLogger) Warn(format string, args ...interface{}) {
	f.log(WarningLevel, format, args...)
}

//Info 方法
func (f *FileLogger) Info(format string, args ...interface{}) {
	f.log(InfoLevel, format, args...)
}

//Error 方法
func (f *FileLogger) Error(format string, args ...interface{}) {
	f.log(ErrorLevel, format, args...)
}

//Fatal 方法
func (f *FileLogger) Fatal(format string, args ...interface{}) {
	f.log(FatalLevel, format, args...)
}

//Close 方法
func (f *FileLogger) Close() {
	f.file.Close()
}
