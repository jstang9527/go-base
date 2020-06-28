package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

//FileLogger 往文件里面写日志
type FileLogger struct {
	Level       LogLevel
	filePath    string
	fileName    string
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
	logChan     chan *logMsg
}

type logMsg struct {
	level     LogLevel
	msg       string
	funcName  string
	fileName  string
	timestamp string
	line      int
}

// NewFileLogger 构造函数
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := parseLevelString(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
		logChan:     make(chan *logMsg, 50000),
	}
	err = fl.initFile() //打开目的日志文件
	if err != nil {
		panic(err)
	}
	return fl
}

func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("打开日志文件失败了, err:", err)
		return err
	}
	errfileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("打开日志文件失败了, err:", err)
		return err
	}
	f.fileObj = fileObj
	f.errFileObj = errfileObj
	//开启一个后台goroutine写日志

	go f.wirteLogByChannel()

	return nil
}

//判定错误等级是否达到客户设定值以上
func (f *FileLogger) enable(LogLevel LogLevel) bool {
	return LogLevel >= f.Level
}

func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("获取文件信息出错, err:", err)
		return false
	}
	return fileInfo.Size() > f.maxFileSize
}

func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	logTime := time.Now().Format("20060102150405")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("获取文件信息出错, err:", err)
		return nil, err
	}
	logName := path.Join(f.filePath, fileInfo.Name())
	newLogName := fmt.Sprintf("%s.bak-%s", logName, logTime)
	//1.关闭当前的日志文件
	file.Close()
	//2.安装时间格式+后缀bak进行备份
	os.Rename(logName, newLogName)
	//3.打开一个新的日志文件
	fileobj, err := os.OpenFile(logName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("新建日志文件出错,无法打开信息日志文件,err:", err)
		return nil, err
	}
	return fileobj, nil
}

func (f *FileLogger) wirteLogByChannel() {
	for {
		if f.checkSize(f.fileObj) { //如果文件太大,则进行切割文件
			newFile, err := f.splitFile(f.fileObj)
			if err != nil {
				return
			}
			f.fileObj = newFile //将打开的文件对象赋值给f.fileObj
		}

		select {
		case logTmp := <-f.logChan:
			logInfo := fmt.Sprintf("[%s] [%s] [%s:%s:%d] %v\n", logTmp.timestamp, getLogLevel(logTmp.level), logTmp.fileName, logTmp.funcName, logTmp.line, logTmp.msg)
			fmt.Fprintf(f.fileObj, logInfo)

			//将Error级别以上日志写入错误文件
			if logTmp.level >= ERROR {
				if f.checkSize(f.errFileObj) { //如果错误文件太大则进行切割
					newErrFile, err := f.splitFile(f.errFileObj)
					if err != nil {
						return
					}
					f.errFileObj = newErrFile //将新的错误文件对象赋值给f.errfileObj
				}
				fmt.Fprintf(f.errFileObj, logInfo)
			}
		default:
			//取不到日志，让出CPU
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func (f *FileLogger) log(lv LogLevel, format string, v ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, v...)
		timeObj := time.Now()
		funcName, funcFile, fileLine := getInfo(3)
		// 造日志对象, 发生到通道
		logTmp := &logMsg{
			level:     lv,
			msg:       msg,
			funcName:  funcName,
			fileName:  funcFile,
			timestamp: timeObj.Format("2006-01-02 15:04:05"),
			line:      fileLine,
		}
		select {
		case f.logChan <- logTmp:
		default: //通道满则丢弃,否则前台程序会卡住
		}

	}
}

//Debug 测试打印
func (f *FileLogger) Debug(format string, v ...interface{}) {
	f.log(DEBUG, format, v...)
}

// Trace 信息
func (f *FileLogger) Trace(format string, v ...interface{}) {
	f.log(TRACE, format, v...)
}

// Info 信息
func (f *FileLogger) Info(format string, v ...interface{}) {
	f.log(INFO, format, v...)
}

// Warning 警告信息
func (f *FileLogger) Warning(format string, v ...interface{}) {
	f.log(WARNING, format, v...)
}

// Error 错误输出
func (f *FileLogger) Error(format string, v ...interface{}) {
	f.log(ERROR, format, v...)
}

// Fatal 奔溃信息
func (f *FileLogger) Fatal(format string, v ...interface{}) {
	f.log(FATAL, format, v...)
}

//Close 关闭文件
func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
