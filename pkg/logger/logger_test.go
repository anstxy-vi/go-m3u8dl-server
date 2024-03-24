package logger

import (
	"errors"
	"testing"
)

func TestLogger(t *testing.T) {
	err := InitDefault("logs/access.log", "logs/error.log", "logs/debug.log", false)
	if err != nil {
		t.Error(err)
	}

	logger.Debug("debug logger")
	logger.Info("info logger")
	logger.Warn("warn logger")
	Error("error logger", errors.New("error is error"))
}
