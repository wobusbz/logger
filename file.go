package logger

import (
	"fmt"
	logs "log"
	"os"
	"path"
	"time"
)

type fileLogger struct {
	level LOGGERLEVELTYPE

	filePath string
	fileName string

	fileInfo  *os.File // 普通日志文件
	errorInfo *os.File // 错误日志文件
	fileSize  int64    // 文件大小

	logChanMsg chan *logMsg
	logMsg     *logMsg
}

func newFileLogger(fileName, filePath string, loggerMax int64) LoggerImpl {
	obj := new(fileLogger)
	obj.fileName = fileName
	obj.filePath = path.Join(filePath, time.Now().Format("20060102"))
	obj.fileSize = loggerMax * 1024
	obj.logChanMsg = make(chan *logMsg, 10000)
	obj.logMsg = newLogMsg()
	obj.Init()
	return obj
}

func (f *fileLogger) createLogFile() {
	fileInfo, err := os.OpenFile(f.getFileInfo(), os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		logs.Fatalf("createLogFile fileInfo: %s", err)
	}

	errorInfo, err := os.OpenFile(f.getErrorInfo(), os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		logs.Fatalf("createLogFile errorInfo: %s", err)
	}
	f.fileInfo = fileInfo
	f.errorInfo = errorInfo
}

func (f *fileLogger) isMkdir() {
	err := os.MkdirAll(f.filePath, 0755)
	f.createLogFile()
	if err != nil {
		logs.Fatalf("isMkdir: %s\n", err)
	}
}

// 判断当天日期文件目录是否存在
func (f *fileLogger) isDay() bool {
	return path.Base(f.filePath) == time.Now().Format("20060102")
}

// 切割日志文件
func (f *fileLogger) splitFileLog(fileInfo, errorInfo string) {
	if !f.isDay() {
		f.fileInfo = f.rename(f.fileInfo, fileInfo)
		f.errorInfo = f.rename(f.fileInfo, errorInfo)
		f.Init()
	}
	if f.isCheckSize(f.fileInfo) {
		f.fileInfo = f.rename(f.fileInfo, fileInfo)
	}
	if f.isCheckSize(f.errorInfo) {
		f.errorInfo = f.rename(f.errorInfo, errorInfo)
	}
}

// 校验文件大小
func (f *fileLogger) isCheckSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		logs.Fatalf("isCheckSize: %s", err)
		return false
	}
	fmt.Println(fileInfo.Size(), f.fileSize)
	return fileInfo.Size() >= f.fileSize
}

// 写文件
func (f *fileLogger) writeLogFile() {
	for {
		logMsg := <-f.logChanMsg
		levelRank, err := getLevelString(logMsg.Level)
		if err != nil {
			logs.Fatalf("writeLogFile: %s\n", err)
		}
		f.splitFileLog(f.getFileInfo(), f.getErrorInfo())
		if logMsg.Level >= ERROR {
			fmt.Fprintf(f.errorInfo, "%s - [%s] - [%s:%d] - %s\n", logMsg.timeDate, levelRank, logMsg.fileName, logMsg.funcLineNo, logMsg.msg)
		} else {
			fmt.Fprintf(f.fileInfo, "%s - [%s] - [%s:%d] - %s\n", logMsg.timeDate, levelRank, logMsg.fileName, logMsg.funcLineNo, logMsg.msg)
		}
	}
}

// 备份
func (f *fileLogger) rename(file *os.File, oldPath string) *os.File {
	file.Close()
	oldFileName := path.Base(oldPath)
	backFileName := fmt.Sprintf("%s_%s", oldFileName, time.Now().Format("200601021504"))
	err := os.Rename(oldPath, path.Join(f.filePath, backFileName))
	if err != nil {
		logs.Fatalf("rename: %s\n", err)
	}
	newFileName, err := os.OpenFile(oldPath, os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		logs.Fatalf("rename newFileName: %s", err)
	}
	return newFileName
}

// 获取完整的fileInfo文件名
func (f *fileLogger) getFileInfo() string {
	return path.Join(f.filePath, fmt.Sprintf("%s_info.log", f.fileName))
}

// 获取完整的errorInfo文件名
func (f *fileLogger) getErrorInfo() string {
	return path.Join(f.filePath, fmt.Sprintf("%s_error.log", f.fileName))
}

func (f *fileLogger) Init() {
	if !exist(f.filePath) {
		f.isMkdir()
	} else {
		f.createLogFile()
	}
	go f.writeLogFile()
}

func (f *fileLogger) SetLevel(level LOGGERLEVELTYPE) {
	if f.level > level || f.level < level {
		f.level = DEBUG
	}
	f.level = level
}

func (f *fileLogger) Write(level LOGGERLEVELTYPE, format string, args ...interface{}) {
	logMsg := f.logMsg.writeLogMsg(level, format, args...)
	select {
	case f.logChanMsg <- logMsg:
	default:
	}
}

func (f *fileLogger) Debug(format string, args ...interface{}) {
	f.Write(DEBUG, format, args...)
}

func (f *fileLogger) Trace(format string, args ...interface{}) {
	f.Write(TRACE, format, args...)
}

func (f *fileLogger) Info(format string, args ...interface{}) {
	f.Write(INFO, format, args...)
}

func (f *fileLogger) Warn(format string, args ...interface{}) {
	f.Write(WARN, format, args...)
}

func (f *fileLogger) Error(format string, args ...interface{}) {
	f.Write(ERROR, format, args...)
}

func (f *fileLogger) Fatal(format string, args ...interface{}) {
	f.Write(FATAL, format, args...)
}

func (f *fileLogger) Close() {

}
