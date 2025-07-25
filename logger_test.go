package golog

import (
	"testing"
)

func TestLog(t *testing.T) {
	logg := New(WithLevel(LevelInfo))
	logg.AddHandler(&Stderr)

	logg.Warn("this is %v", 1)
	logg.Info("this is %v", 1)
	logg.Debug("this is %v", 1)

}
