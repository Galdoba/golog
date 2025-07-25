package golog

import (
	"fmt"
	"time"
)

type Formatter interface {
	Format(LogLevel, string, ...any) string
}

type PlainTextFormatter struct{}

func (pt *PlainTextFormatter) Format(level LogLevel, format string, args ...any) string {
	timestamp := time.Now().Format(time.DateTime + ".000")
	return fmt.Sprintf("[%s] [%s]: %s", timestamp, levelToString(level), fmt.Sprintf(format, args...))
}
