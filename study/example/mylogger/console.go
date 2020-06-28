package mylogger

import (
	"fmt"
	"time"
)

//ConsoleLogger 结构体
type ConsoleLogger struct {
	Level LogLevel
}

// NewConsolelog 结构体构造函数
func NewConsolelog(levelStr string) ConsoleLogger {
	level, err := parseLevelString(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}
}

//判定错误等级是否达到客户设定值以上
func (c ConsoleLogger) enable(LogLevel LogLevel) bool {
	return LogLevel >= c.Level
}

func (c ConsoleLogger) log(lv LogLevel, format string, v ...interface{}) {
	if c.enable(lv) {
		msg := fmt.Sprintf(format, v...)
		timeObj := time.Now()
		funcName, FuncFile, fileLine := getInfo(3)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %v\n", timeObj.Format("2006-01-02 15:04:05 MST"), getLogLevel(lv), FuncFile, funcName, fileLine, msg)
	}
}

//Debug 测试打印
func (c ConsoleLogger) Debug(format string, v ...interface{}) {
	c.log(DEBUG, format, v...)
}

// Trace 信息
func (c ConsoleLogger) Trace(format string, v ...interface{}) {
	c.log(TRACE, format, v...)
}

// Info 信息
func (c ConsoleLogger) Info(format string, v ...interface{}) {
	c.log(INFO, format, v...)
}

// Warning 警告信息
func (c ConsoleLogger) Warning(format string, v ...interface{}) {
	c.log(WARNING, format, v...)
}

// Error 错误输出
func (c ConsoleLogger) Error(format string, v ...interface{}) {
	c.log(ERROR, format, v...)
}

// Fatal 奔溃信息
func (c ConsoleLogger) Fatal(format string, v ...interface{}) {
	c.log(FATAL, format, v...)
}
