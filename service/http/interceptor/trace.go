package interceptor

import (
	"net/http"

	"github.com/te2ka/golang-interceptor-example/utils/tracer"
)

func WithTrace() Opt {
	return func(o *opt) {
		o.registerInterceptor(traceInterceptor)
	}
}

func traceInterceptor(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h(w, r.WithContext(tracer.WithTraceID(r.Context())))
	}
}
