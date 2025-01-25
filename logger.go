package log

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/lz-go/config"
)

// NewLogger creates a new opinionated Logger instance.
// The option could be provided for customization.
func NewLogger(opts ...Option) Logger {
	i := &loggerImpl{
		ctx:    context.Background(),
		fields: make([]zap.Field, 0),
	}

	for _, opt := range opts {
		opt(i)
	}

	i.logger = i.prepareLogger()

	return i
}

type loggerImpl struct {
	ctx    context.Context
	fields []zap.Field
	logger *zap.Logger
}

// WithField adds structured context to the current Logger instance.
func (i *loggerImpl) WithField(key string, value any) {
	i.fields = append(i.fields, zap.Any(key, value))
}

// Debug logs a message at zap.DebugLevel. The message includes any fields
// accumulated on the Logger instance.
func (i *loggerImpl) Debug(msg string) {
	i.logger.Debug(msg, i.fields...)
}

// Debugf logs a formated message at zap.DebugLevel. The message includes any
// fields accumulated on the Logger instance.
func (i *loggerImpl) Debugf(format string, a ...any) {
	i.logger.Debug(fmt.Sprintf(format, a...))
}

// Info logs a message at zap.InfoLevel. The message includes any fields
// accumulated on the Logger instance.
func (i *loggerImpl) Info(msg string) {
	i.logger.Info(msg)
}

// Infof logs a formated message at zap.InfoLevel. The message includes any
// fields accumulated on the Logger instance.
func (i *loggerImpl) Infof(format string, a ...any) {
	i.logger.Info(fmt.Sprintf(format, a...))
}

// Warn logs a message at zap.WarnLevel. The message includes any fields
// accumulated on the Logger instance.
func (i *loggerImpl) Warn(msg string) {
	i.logger.Warn(msg)
}

// Warnf logs a formated message at zap.WarnLevel. The message includes any
// fields accumulated on the Logger instance.
func (i *loggerImpl) Warnf(format string, a ...any) {
	i.logger.Warn(fmt.Sprintf(format, a...))
}

// Error logs a message at zap.ErrorLevel. The message includes any fields
// accumulated on the Logger instance.
func (i *loggerImpl) Error(msg string) {
	i.logger.Error(msg)
}

// Errorf logs a formated message at zap.ErrorLevel. The message includes any
// fields accumulated on the Logger instance.
func (i *loggerImpl) Errorf(format string, a ...any) {
	i.logger.Error(fmt.Sprintf(format, a...))
}

// Panic logs a message at zap.PanicLevel. The message includes any fields
// accumulated on the Logger instance.
func (i *loggerImpl) Panic(msg string) {
	i.logger.Panic(msg)
}

// Panicf logs a formated message at zap.PanicLevel. The message includes any
// fields accumulated on the Logger instance.
func (i *loggerImpl) Panicf(format string, a ...any) {
	i.logger.Panic(fmt.Sprintf(format, a...))
}

// Fatal logs a message at zap.FatalLevel. The message includes any fields
// accumulated on the Logger instance.
func (i *loggerImpl) Fatal(msg string) {
	i.logger.Fatal(msg)
}

// Fatalf logs a formated message at zap.FatalLevel. The message includes any
// fields accumulated on the Logger instance.
func (i *loggerImpl) Fatalf(format string, a ...any) {
	i.logger.Fatal(fmt.Sprintf(format, a...))
}

// prepareLogger prepare the zap.Logger instance with the preloaded configuration.
// Refer to the `README.md` document to understand which config will be used.
func (i *loggerImpl) prepareLogger() *zap.Logger {
	cfg := zap.NewProductionConfig()
	if config.GetStringOr("app.env", "local") == "local" {
		cfg = zap.NewDevelopmentConfig()
	}

	switch config.GetStringOr("log.level", "debug") {
	case "debug":
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		cfg.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		cfg.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	default:
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	switch config.GetStringOr("log.encoding", "json") {
	case "console":
		cfg.Encoding = "console"
	case "json":
		cfg.Encoding = "json"
	default:
		cfg.Encoding = "json"
	}

	logger, _ := cfg.Build()

	if v, ok := i.ctx.Value(CorrelationIdContextKey).(string); ok {
		i.fields = append(i.fields, zap.String(CorrelationIdLogKey, v))
	}

	return logger.With(i.fields...)
}

// Option supports customize the Logger instance while initializing.
type Option func(*loggerImpl)

// WithContext adds the context into the Logger instance, it will allow the
// logger instance to include the correlation ID if attached.
func WithContext(ctx context.Context) Option {
	return func(i *loggerImpl) {
		i.ctx = ctx
	}
}

// WithField adds the key-value field into the Logger instance.
func WithField(key string, value any) Option {
	return func(i *loggerImpl) {
		i.fields = append(i.fields, zap.Any(key, value))
	}
}

func WithFunc(value string) Option {
	return func(i *loggerImpl) {
		i.fields = append(i.fields, zap.Any(FuncLogKey, value))
	}
}
