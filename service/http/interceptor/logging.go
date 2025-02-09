package interceptor

import (
	"net/http"

	"github.com/te2ka/golang-interceptor-example/utils/log"
	"github.com/te2ka/golang-interceptor-example/utils/tracer"
)

func WithLogging() Opt {
	return func(o *opt) {
		o.interceptors = append(o.interceptors, loggingInterceptor)
	}
}

func loggingInterceptor(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := tracer.WithTraceID(r.Context())
		r = r.WithContext(ctx)

		log.Logger.InfoContext(ctx, "start")
		h(w, r)
		log.Logger.InfoContext(ctx, "end")
	}
}
