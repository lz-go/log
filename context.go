package log

import (
	"context"

	"github.com/google/uuid"
)

func CorrelationContext(ctx context.Context) context.Context {
	if _, ok := ctx.Value(CorrelationIdContextKey).(string); ok {
		return ctx
	}

	return context.WithValue(ctx, CorrelationIdContextKey, uuid.New().String())
}
