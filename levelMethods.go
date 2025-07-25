package golog

func (l *Logger) Trace(msg string, args ...interface{}) {
	l.log(LevelTrace, msg, args...)
}

func (l *Logger) Debug(msg string, args ...interface{}) {
	l.log(LevelDebug, msg, args...)
}

func (l *Logger) Info(msg string, args ...interface{}) {
	l.log(LevelInfo, msg, args...)
}

func (l *Logger) Warn(msg string, args ...interface{}) {
	l.log(LevelWarn, msg, args...)
}

func (l *Logger) Notice(msg string, args ...interface{}) {
	l.log(LevelNotice, msg, args...)
}

func (l *Logger) Error(msg string, args ...interface{}) {
	l.log(LevelError, msg, args...)
}

func (l *Logger) Critical(msg string, args ...interface{}) {
	l.log(LevelCritical, msg, args...)
}

func (l *Logger) Alert(msg string, args ...interface{}) {
	l.log(LevelAlert, msg, args...)
}

func (l *Logger) Fatal(msg string, args ...interface{}) {
	l.log(LevelFatal, msg, args...)
}
func (l *Logger) Emergency(msg string, args ...interface{}) {
	l.log(LevelEmergency, msg, args...)
}

var levelNames = map[LogLevel]string{
	LevelTrace:     "trace",
	LevelDebug:     "debug",
	LevelInfo:      "info",
	LevelNotice:    "notice",
	LevelWarn:      "warning",
	LevelError:     "error",
	LevelCritical:  "critical",
	LevelAlert:     "alert",
	LevelFatal:     "fatal",
	LevelEmergency: "emergency",
}

func levelToString(l LogLevel) string {
	return levelNames[l]
}
