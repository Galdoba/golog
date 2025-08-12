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
		switch handler.destination {
		case Stderr, Stdout:
			fmt.Fprintln(handler.descriptor, msg)
		default:
			f, _ := os.OpenFile(handler.destination, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
			defer f.Close()
			fmt.Fprintln(f, msg)

		}

	}
	if mustShutdown {
		os.Exit(1)
	}

}
