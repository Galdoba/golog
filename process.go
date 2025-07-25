package golog

import (
	"fmt"
)

func (l *Logger) log(level LogLevel, msg string, args ...any) {
	for _, handler := range l.handlers {
		if level < l.level {
			return
		}

		l.mu.Lock()
		defer l.mu.Unlock()

		msg := handler.formatter.Format(level, msg, args...)
		fmt.Fprintln(handler.writer, msg)

	}

}

// func (l *Logger) formatMessage(level LogLevel, msg string, args ...any) string {
// 	levelName := levelToString(level)

// 	if l.useColors {
// 		levelName = l.colorize(level, levelName)
// 	}

// 	return fmt.Sprintf("[%s] %s", levelName, fmt.Sprintf(msg, args...))
// }

// func (l *Logger) colorize(level LogLevel, text string) string {
// 	colors := map[LogLevel]string{
// 		LevelFatal: "\033[1;31m", // Красный
// 		LevelError: "\033[31m",   // Красный
// 		LevelWarn:  "\033[33m",   // Желтый
// 		LevelInfo:  "\033[36m",   // Голубой
// 		LevelDebug: "\033[35m",   // Пурпурный
// 		LevelTrace: "\033[37m",   // Серый
// 	}

// 	if color, ok := colors[level]; ok {
// 		return color + text + "\033[0m" // Сброс цвета
// 	}
// 	return text
// }
