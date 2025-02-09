package interceptor

import (
	"net/http"
	"sync"
)

var (
	wg sync.WaitGroup
)

func WithCountConnection() Opt {
	return func(o *opt) {
		o.interceptors = append(o.interceptors, countConnectionInterceptor)
	}
}

func countConnectionInterceptor(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wg.Add(1)
		h(w, r)
		wg.Done()
	}
}
