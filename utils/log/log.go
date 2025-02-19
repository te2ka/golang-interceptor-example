package log

import (
	"context"
	"log/slog"
	"os"

	"github.com/te2ka/golang-interceptor-example/utils/tracer"
)

func init() {
	handler := &logHandler{
		slog.NewJSONHandler(os.Stdout, nil),
	}
	logger := slog.New(handler)
	slog.SetDefault(logger)
}

type logHandler struct {
	slog.Handler
}

func (h *logHandler) Handle(ctx context.Context, r slog.Record) error {
	r.AddAttrs(slog.String("trace_id", tracer.GetTraceID(ctx)))
	return h.Handler.Handle(ctx, r)
}
