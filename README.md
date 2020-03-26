# logger
	// filename  文件名
	// filePath  文件路径
	// true  输入终端
	// false  输入到文件夹
	logger := NewLogger()
	logger.FileName = "abc"
	logger.FilePath = "abc"
	logger.SetLoggerMax(300)
	logger.SetConsole(false)
	CustomLogger(logger)
	
	Debug("this is Debug: %s", "wuhuarou")
	Info("this is Info: %s", "wuhuarou")
	Warn("this is Warn: %s", "wuhuarou")
	Error("this is Error: %s", "wuhuarou")
	Fatal("this is Fatal: %s", "wuhuarou")
	time.Sleep(time.Second)


