package log

import (
	"log/slog"
)

var (
	def = NewLogger(slog.Default())
)

type Logger interface {
	Debug(message string, args ...any)
	Info(message string, args ...any)
	Warn(message string, args ...any)
	Error(message string, args ...any)
	With(args ...any) Logger
}

type logger struct {
	*slog.Logger
}

func (l *logger) With(args ...any) Logger {
	return NewLogger(l.Logger.With(args...))
}

func NewLogger(l *slog.Logger) Logger {
	return &logger{l}
}

func Debug(message string, args ...any) {
	def.Debug(message, args...)
}

func Info(message string, args ...any) {
	def.Info(message, args...)
}

func Warn(message string, args ...any) {
	def.Warn(message, args...)
}

func Error(message string, args ...any) {
	def.Error(message, args...)
}

func With(args ...any) Logger {
	return def.With(args...)
}

func SetDefault(l Logger) {
	def = l
}
