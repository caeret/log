package log

import (
	"log/slog"
)

var _ Logger = (*logger)(nil)

type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
	With(args ...any) Logger
}

func New(l *slog.Logger) Logger {
	return &logger{l: l}
}

type logger struct {
	l *slog.Logger
}

func (l *logger) Debug(msg string, args ...any) {
	l.l.Debug(msg, args...)
}

func (l *logger) Info(msg string, args ...any) {
	l.l.Info(msg, args...)
}

func (l *logger) Warn(msg string, args ...any) {
	l.l.Warn(msg, args...)
}

func (l *logger) Error(msg string, args ...any) {
	l.l.Error(msg, args...)
}

func (l *logger) With(args ...any) Logger {
	return New(l.l.With(args...))
}
