package logger

var (
	log          LoggerImpl
	DEFAULTLEVEL LOGGERLEVELTYPE = 0
	LOG_FILE_MAX int64           = 10
)

type Logger struct {
	console  bool
	file     bool
	defaults bool
}

func NewLogger(fileName, filePath string, b bool) {
	if b {
		log = newConsoleLog()
	} else {
		log = newFileLogger(fileName, filePath)
	}
}

func (l *Logger) SetConsole(b bool) {
	l.console = b
}

func (l *Logger) SetFile(b bool) {
	l.file = b
}

func Debug(format string, args ...interface{}) {
	log.Debug(format, args...)
}

func Trace(format string, args ...interface{}) {
	log.Trace(format, args...)
}

func Info(format string, args ...interface{}) {
	log.Info(format, args...)
}

func Warn(format string, args ...interface{}) {
	log.Warn(format, args...)
}

func Error(format string, args ...interface{}) {
	log.Error(format, args...)
}

func Fatal(format string, args ...interface{}) {
	log.Fatal(format, args...)
}
