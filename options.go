package golog

type LoggerOption func(*Logger)

func WithLevel(level LogLevel) LoggerOption {
	return func(l *Logger) {
		l.level = level
	}
}

// func WithOutput(writers ...io.Writer) LoggerOption {
// 	return func(l *Logger) {
// 		l.outputs = append(l.outputs, writers...)
// 	}
// }

// func WithPrefix(prefix string) LoggerOption {
// 	return func(l *Logger) {
// 		l.prefix = prefix
// 	}
// }

// func WithColor(enabled bool) LoggerOption {
// 	return func(l *Logger) {
// 		l.useColors = enabled
// 	}
// }

// func WithFlags(flags int) LoggerOption {
// 	return func(l *Logger) {
// 		l.flags = flags
// 	}
// }
