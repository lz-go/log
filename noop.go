package log

// NewNoopLogger create a no operation instance of Logger.
func NewNoopLogger() Logger {
	return &noopLoggerImpl{}
}

type noopLoggerImpl struct{}

func (noopLoggerImpl) WithField(key string, value interface{}) {}

func (noopLoggerImpl) Debug(msg string) {}

func (noopLoggerImpl) Debugf(format string, a ...any) {}

func (noopLoggerImpl) Info(msg string) {}

func (noopLoggerImpl) Infof(format string, a ...any) {}

func (noopLoggerImpl) Warn(msg string) {}

func (noopLoggerImpl) Warnf(format string, a ...any) {}

func (noopLoggerImpl) Error(msg string) {}

func (noopLoggerImpl) Errorf(format string, a ...any) {}

func (noopLoggerImpl) Panic(msg string) {}

func (noopLoggerImpl) Panicf(format string, a ...any) {}

func (noopLoggerImpl) Fatal(msg string) {}

func (noopLoggerImpl) Fatalf(format string, a ...any) {}
