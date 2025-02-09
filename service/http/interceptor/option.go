package interceptor

import (
	"net/http"
)

type opt struct {
	interceptors []Interceptor
}

type Opt func(*opt)

func ApplyInterceptors(handler http.HandlerFunc, opts ...Opt) http.HandlerFunc {
	o := &opt{}
	for _, option := range opts {
		option(o)
	}

	h := handler
	for _, interceptor := range o.interceptors {
		h = interceptor(h)
	}

	return h
}
