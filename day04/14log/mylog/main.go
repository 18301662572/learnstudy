package mylog

//自定义一个日志库，实现日志记录的功能

//日志分级别
// debug trace info warn error cirtal ...

const (
	DEBUG = iota
	TRACE
	INFO
	WARN
	ERROR
	CIRTAL
)

func getlevelStr(level int) string {
	switch level {
	case 0:
		return "DEBUG"
	case 1:
		return "TRACE"
	case 2:
		return "INFO"
	case 3:
		return "WARN"
	case 4:
		return "ERROR"
	case 5:
		return "CIRTAL"
	default:
		return "DEBUG"
	}
}
