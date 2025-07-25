package golog

import (
	"fmt"
)

func (l *Logger) log(level LogLevel, msg string, args ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()
	mustShutdown := false
	for _, handler := range l.handlers {
		if level < l.level {
			return
		}
		if level >= LevelFatal {
			mustShutdown = true
		}

		msg := handler.formatter.Format(level, msg, args...)
		fmt.Fprintln(handler.writer, msg)

	}
	if mustShutdown {
		fmt.Printf("shutdown now: %v\n", level)
	}

}
