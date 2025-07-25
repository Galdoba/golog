package golog

import (
	"io"
	"os"
)

type MessageHandler struct {
	writer    io.Writer
	formatter Formatter
	minLevel  LogLevel
}

var Stderr MessageHandler = MessageHandler{
	writer:    os.Stderr,
	formatter: &PlainTextFormatter{},
	minLevel:  LevelInfo,
}

func (l *Logger) AddHandler(handler *MessageHandler) {
	l.handlers = append(l.handlers, handler)
}

func NewHandler(writer io.Writer, level LogLevel, formatter Formatter) *MessageHandler {
	return &MessageHandler{
		writer:    writer,
		formatter: formatter,
		minLevel:  level,
	}
}
