package main

import (
	"time"

	"github.com/jstang007/gateway_demo/study/example/mylogger"
)

var logger mylogger.Logger

func publicLog() {
	logger = mylogger.NewConsolelog("Debug")
	logger = mylogger.NewFileLogger("debug", "./", "knho.log", 1*1024)
}

func consoleLog() {
	loger := mylogger.NewConsolelog("Debug")
	loger.Debug("我是%v一个debug%v", true, [...]int{1, 2, 3})
	loger.Trace("我是一个Trace%v", true)
	loger.Info("我是一个信息")
	loger.Error("我是错误Error")
	loger.Warning("我是警告")
	loger.Fatal("我奔溃了")
}

func fileLog() {
	loger := mylogger.NewFileLogger("debug", "./", "knho.log", 5*1024) //1024字节=1kb,1024kb=1M
	for {
		loger.Debug("我是%v一个debug%v", true, [...]int{1, 2, 3})
		loger.Trace("我是一个Trace%v", true)
		loger.Info("我是一个信息")
		loger.Debug("我是%v一个debug%v", true, [...]int{1, 2, 3})
		loger.Trace("我是一个Trace%v", true)
		loger.Info("我是一个信息")
		loger.Error("我是错误Error0000001")
		loger.Warning("我是警告")
		loger.Error("我是错误Error0000002")
		loger.Fatal("我奔溃了1我奔溃了1我奔溃了1我奔溃了1我奔溃了1我奔溃了1")
		loger.Fatal("我奔溃了2我奔溃了2我奔溃了2我奔溃了2我奔溃了2我奔溃了2")
		time.Sleep(time.Second * 3)
	}
}

func main() {
	fileLog()

	//有个问题,每个包都往里面写数据时(同时打开同一个文件写入)，使用管道即可
}
