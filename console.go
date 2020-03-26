package logger

import (
	"fmt"
	"os"
)

type consoleLog struct {
	level  LOGGERLEVELTYPE
	logMsg *logMsg
}

func newConsoleLog() *consoleLog {
	obj := &consoleLog{}
	obj.Init()
	return obj
}

func (c *consoleLog) Init() {
	c.logMsg = newLogMsg()
}

func (c *consoleLog) SetLevel(level LOGGERLEVELTYPE) {

}

func (c *consoleLog) Write(level LOGGERLEVELTYPE, format string, args ...interface{}) {
	logMsg := c.logMsg.writeLogMsg(level, format, args...)
	if levelRank, err := getLevelString(level); err == nil {
		fmt.Fprintf(os.Stdout, "%s - [%s] - [%s:%d] - %s\n", logMsg.timeDate, levelRank, logMsg.fileName, logMsg.funcLineNo, logMsg.msg)
	}
}

func (c *consoleLog) Debug(format string, args ...interface{}) {
	c.Write(DEBUG, format, args...)
}

func (c *consoleLog) Trace(format string, args ...interface{}) {
	c.Write(TRACE, format, args...)
}

func (c *consoleLog) Info(format string, args ...interface{}) {
	c.Write(INFO, format, args...)
}

func (c *consoleLog) Warn(format string, args ...interface{}) {
	c.Write(WARN, format, args...)
}

func (c *consoleLog) Error(format string, args ...interface{}) {
	c.Write(ERROR, format, args...)
}

func (c *consoleLog) Fatal(format string, args ...interface{}) {
	c.Write(FATAL, format, args...)
}

func (c *consoleLog) Close() {

}
