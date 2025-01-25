package log

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lz-go/config"
)

func TestLogger(t *testing.T) {
	t.Run("TestNewLogger", func(t *testing.T) {
		t.Run("returns a new Logger instance", func(t *testing.T) {
			config.Reset()

			logger := NewLogger()
			assert.NotNil(t, logger)
		})

		t.Run("returns a new Logger instance with options", func(t *testing.T) {
			config.Reset()

			ctx := context.WithValue(context.Background(), "cxid", "abcdef0123456789")

			logger := NewLogger(
				WithContext(ctx),
				WithFunc("TestNewLogger"),
				WithField("key_1", "value"),
				WithField("key_2", "value"),
			)
			assert.NotNil(t, logger)
		})

		t.Run("returns a new Logger instance with debug level", func(t *testing.T) {
			config.Reset()
			config.Set("log.level", "debug")

			logger := NewLogger()
			assert.NotNil(t, logger)
		})

		t.Run("returns a new Logger instance with info level", func(t *testing.T) {
			config.Reset()
			config.Set("log.level", "info")

			logger := NewLogger()
			assert.NotNil(t, logger)
		})

		t.Run("returns a new Logger instance with warn level", func(t *testing.T) {
			config.Reset()
			config.Set("log.level", "warn")

			logger := NewLogger()
			assert.NotNil(t, logger)
		})

		t.Run("returns a new Logger instance with error level", func(t *testing.T) {
			config.Reset()
			config.Set("log.level", "error")

			logger := NewLogger()
			assert.NotNil(t, logger)
		})

		t.Run("returns a new Logger instance with debug level when the config is invalid", func(t *testing.T) {
			config.Reset()
			config.Set("log.level", "invalid")

			logger := NewLogger()
			assert.NotNil(t, logger)
		})

		t.Run("returns a new Logger instance with console encoding", func(t *testing.T) {
			config.Reset()
			config.Set("log.encoding", "console")

			logger := NewLogger()
			assert.NotNil(t, logger)
		})

		t.Run("returns a new Logger instance with json encoding", func(t *testing.T) {
			config.Reset()
			config.Set("log.encoding", "json")

			logger := NewLogger()
			assert.NotNil(t, logger)
		})

		t.Run("returns a new Logger instance with json encoding when the config is invalid", func(t *testing.T) {
			config.Reset()
			config.Set("log.encoding", "invalid")

			logger := NewLogger()
			assert.NotNil(t, logger)
		})
	})

	t.Run("TestLogger", func(t *testing.T) {
		config.Reset()
		config.Set("app.env", "local")

		logger := NewLogger()

		t.Run("WithField", func(t *testing.T) {
			assert.NotPanics(t, func() {
				logger.WithField("key", "value")
			})
		})

		t.Run("Debug", func(t *testing.T) {
			assert.NotPanics(t, func() {
				logger.Debug("Hello, world!")
			})
		})

		t.Run("Debugf", func(t *testing.T) {
			assert.NotPanics(t, func() {
				logger.Debugf("Hello, %s!", "world")
			})
		})

		t.Run("Info", func(t *testing.T) {
			assert.NotPanics(t, func() {
				logger.Info("Hello, world!")
			})
		})

		t.Run("Infof", func(t *testing.T) {
			assert.NotPanics(t, func() {
				logger.Infof("Hello, %s!", "world")
			})
		})

		t.Run("Warn", func(t *testing.T) {
			assert.NotPanics(t, func() {
				logger.Warn("Hello, world!")
			})
		})

		t.Run("Warnf", func(t *testing.T) {
			assert.NotPanics(t, func() {
				logger.Warnf("Hello, %s!", "world")
			})
		})

		t.Run("Error", func(t *testing.T) {
			assert.NotPanics(t, func() {
				logger.Error("Hello, world!")
			})
		})

		t.Run("Errorf", func(t *testing.T) {
			assert.NotPanics(t, func() {
				logger.Errorf("Hello, %s!", "world")
			})
		})

		t.Run("Panic", func(t *testing.T) {
			assert.Panics(t, func() {
				logger.Panic("Hello, world!")
			})
		})

		t.Run("Panicf", func(t *testing.T) {
			assert.Panics(t, func() {
				logger.Panicf("Hello, %s!", "world")
			})
		})
	})
}
