package logger

import (
	"testing"
	"time"
)

func TestNewLogger(t *testing.T) {
	NewLogger("abc", "abc", true)
	for {
		Debug("this is Debug: %s", "wuhuarou")
		Info("this is Info: %s", "wuhuarou")
		Warn("this is Warn: %s", "wuhuarou")
		Error("this is Error: %s", "wuhuarou")
		Fatal("this is Fatal: %s", "wuhuarou")
		time.Sleep(time.Second)
	}
}
