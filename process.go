package golog

import (
	"fmt"
	"os"
)

func (l *Logger) log(level LogLevel, msg string, args ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()
	mustShutdown := false
	for _, handler := range l.handlers {
		if level < handler.minLevel {
			continue
		}
		if level >= LevelFatal {
			mustShutdown = true
		}

		msg := handler.formatter.Format(level, msg, args...)
		fmt.Fprintln(handler.writer, msg)

	}
	if mustShutdown {
		os.Exit(1)
	}

}
