package logger

import (
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"time"
)

type LOGGERLEVELTYPE int

const (
	DEBUG LOGGERLEVELTYPE = iota
	INFO
	TRACE
	WARN
	ERROR
	FATAL
)

type logMsg struct {
	Level      LOGGERLEVELTYPE
	fileName   string
	funcName   string
	funcLineNo LOGGERLEVELTYPE
	timeDate   string
	msg        string
}

func newLogMsg() *logMsg {
	return new(logMsg)
}

func (l *logMsg) writeLogMsg(level LOGGERLEVELTYPE, format string, args ...interface{}) *logMsg {
	timeDate := time.Now().Format("2006-01-02 15:04:05.00")
	fileName, funcName, funcLine := getFuncInfo()
	objLogMsg := newLogMsg()
	objLogMsg.Level = level
	objLogMsg.fileName = fileName
	objLogMsg.funcName = funcName
	objLogMsg.funcLineNo = funcLine
	objLogMsg.timeDate = timeDate
	objLogMsg.msg = fmt.Sprintf(format, args...)
	return objLogMsg
}

type LoggerImpl interface {
	Init()
	SetLevel(level LOGGERLEVELTYPE)
	Write(level LOGGERLEVELTYPE, format string, args ...interface{})
	Debug(format string, args ...interface{})
	Trace(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	Close()
}

func getLevelString(level LOGGERLEVELTYPE) (string, error) {
	switch level {
	case DEBUG:
		return "DEBUG", nil
	case TRACE:
		return "TRACE", nil
	case INFO:
		return "INFO", nil
	case WARN:
		return "WARN", nil
	case ERROR:
		return "ERROR", nil
	case FATAL:
		return "FATAL", nil
	default:
		err := errors.New("log rank not")
		return "NOT FOUND", err
	}
}

func getFuncInfo() (fileName, funcName string, funcLineNo LOGGERLEVELTYPE) {
	pc, file, line, ok := runtime.Caller(4)
	if ok {
		fileName = path.Base(file)
		funcName = path.Base(runtime.FuncForPC(pc).Name())
		funcLineNo = LOGGERLEVELTYPE(line)
	}
	return
}

func exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
