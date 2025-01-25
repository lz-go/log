package log

const (
	// CorrelationIdContextKey is the key used to store the correlation ID in a context.
	CorrelationIdContextKey = "cxid"
)

const (
	// CorrelationIdLogKey is the key used to log the correlation ID.
	CorrelationIdLogKey = "cxid"

	// FuncLogKey is the key used to log the function name.
	FuncLogKey = "func"
)

type Logger interface {
	WithField(key string, value interface{})

	Debug(msg string)
	Debugf(format string, a ...any)

	Info(msg string)
	Infof(format string, a ...any)

	Warn(msg string)
	Warnf(format string, a ...any)

	Error(msg string)
	Errorf(format string, a ...any)

	Panic(msg string)
	Panicf(format string, a ...any)

	Fatal(msg string)
	Fatalf(format string, a ...any)
}
