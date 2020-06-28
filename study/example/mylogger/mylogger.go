package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

// LogLevel 定义类型
type LogLevel uint8

// Logger ...
type Logger interface {
	Debug(format string, v ...interface{})
	Trace(format string, v ...interface{})
	Info(format string, v ...interface{})
	Warning(format string, v ...interface{})
	Error(format string, v ...interface{})
	Fatal(format string, v ...interface{})
}

const (
	//UNKNOWN 无效的等级
	UNKNOWN LogLevel = iota
	//DEBUG dev
	DEBUG
	//TRACE x
	TRACE
	//INFO x
	INFO
	// WARNING x
	WARNING
	// ERROR x
	ERROR
	// FATAL x
	FATAL
)

func parseLevelString(s string) (LogLevel, error) {
	switch strings.ToLower(s) {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志等级")
		return UNKNOWN, err
	}
}

//将错误等级转换成string
func getLogLevel(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case TRACE:
		return "TRACE"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

//获取调用日志工具文件函数位置
func getInfo(skip int) (funcName, funcFile string, FileLine int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("runtime.Caller() failed")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	funcFile = path.Base(file)
	FileLine = line
	return
}
