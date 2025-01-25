package log

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCorrelationContext(t *testing.T) {
	t.Run("returns a new context with cxid", func(t *testing.T) {
		ctx := CorrelationContext(context.Background())

		v, ok := ctx.Value("cxid").(string)
		assert.True(t, ok)
		assert.NotEqual(t, v, "")
	})

	t.Run("returns the current context when cxid is already set", func(t *testing.T) {
		ctx := context.WithValue(context.Background(), "cxid", "abcdef0123456789")
		ctx = CorrelationContext(ctx)

		v, ok := ctx.Value("cxid").(string)
		assert.True(t, ok)
		assert.Equal(t, v, "abcdef0123456789")
	})
}
