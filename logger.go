package logger

var (
	log          LoggerImpl
	DEFAULTLEVEL LOGGERLEVELTYPE = 0
	log_file_max int64           = 100
)

type Logger struct {
	Defaults  bool
	LoggerMax int64

	FileName string
	FilePath string
}

func NewLogger() *Logger {
	return new(Logger)

}

// 默认配置
func DefaultLogger() {
	log = newConsoleLog()
}

func CustomLogger(logger *Logger) {
	if logger.Defaults {
		log = newConsoleLog()
	} else {
		log = newFileLogger(logger.FileName, logger.FilePath, logger.LoggerMax)
	}
}

func (l *Logger) SetLoggerMax(loggerMax int64) {
	l.LoggerMax = loggerMax
}

func (l *Logger) SetConsole(b bool) {
	l.Defaults = b
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
