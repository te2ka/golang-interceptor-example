package tracer

import (
	"context"

	"github.com/google/uuid"
)

type traceID struct{}

var traceIDKey = traceID{}

func WithTraceID(ctx context.Context) context.Context {
	return context.WithValue(ctx, traceIDKey, uuid.NewString())
}

func GetTraceID(ctx context.Context) string {
	traceID, _ := ctx.Value(traceIDKey).(string)
	return traceID
}
