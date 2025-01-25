package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoopLogger(t *testing.T) {
	t.Run("NewNoopLogger", func(t *testing.T) {
		assert.NotNil(t, NewNoopLogger())
	})

	logger := NewNoopLogger()

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
		assert.NotPanics(t, func() {
			logger.Panic("Hello, world!")
		})
	})

	t.Run("Panicf", func(t *testing.T) {
		assert.NotPanics(t, func() {
			logger.Panicf("Hello, %s!", "world")
		})
	})

	t.Run("Fatal", func(t *testing.T) {
		assert.NotPanics(t, func() {
			logger.Fatal("Hello, world!")
		})
	})

	t.Run("Fatalf", func(t *testing.T) {
		assert.NotPanics(t, func() {
			logger.Fatalf("Hello, %s!", "world")
		})
	})
}
