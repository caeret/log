package log

var Nop Logger = nopLogger{}

type nopLogger struct{}

func (n nopLogger) Debug(message string, fields ...interface{}) {

}

func (n nopLogger) Info(message string, fields ...interface{}) {

}

func (n nopLogger) Warn(message string, fields ...interface{}) {

}

func (n nopLogger) Error(message string, fields ...interface{}) {

}

func (n nopLogger) With(fields ...interface{}) Logger {
	return n
}
