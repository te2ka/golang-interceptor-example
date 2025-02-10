package interceptor

import (
	"log/slog"
	"net/http"
)

func WithLogging() Opt {
	return func(o *opt) {
		o.registerInterceptor(loggingInterceptor)
	}
}

func loggingInterceptor(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.InfoContext(r.Context(), "start")
		h(w, r)
		slog.InfoContext(r.Context(), "end")
	}
}
