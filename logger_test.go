package golog

import (
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	logg := New(WithLevel(LevelTrace))
	// logg.AddHandler(&Stderr)
	logg.AddHandler(NewHandler(os.Stderr, LevelInfo,
		NewTextFormatter(
			WithTimePrecision(3),
			WithColor(),
		)))
	logg.Emergency("this is %v", "EMERGENCY")
	logg.Fatal("this is %v", "FATAL")
	logg.Alert("this is %v", "ALERT")
	logg.Critical("this is %v", "CRITICAL")
	logg.Error("this is %v", "ERROR")
	logg.Warn("this is %v", "WARN")
	logg.Notice("this is %v", "NOTICE")
	logg.Info("this is %v", "INFO")
	logg.Debug("this is %v", "DEBUG")
	logg.Trace("this is %v", "TRACE")

}
