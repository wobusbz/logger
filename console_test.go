package logger

import (
	"testing"
	"time"
)

func TestNewLogger(t *testing.T) {
	NewLogger("abc", "abc", false)
	for {
		log.Debug("%d", 111)
		time.Sleep(time.Second)
	}
}
