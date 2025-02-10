package server

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/te2ka/golang-interceptor-example/service/http/interceptor"
)

var (
	wg sync.WaitGroup
)

type customHandler struct {
	serveMux *http.ServeMux
}

func (h *customHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	wg.Add(1)

	handler, _ := h.serveMux.Handler(r)
	handlerFunc := interceptor.ApplyInterceptors(handler.ServeHTTP, interceptor.WithTrace(), interceptor.WithLogging())
	handlerFunc(w, r)

	wg.Done()
}

func NewServer(port int) http.Server {
	mux := http.NewServeMux()

	registerRoutes(mux)

	handler := &customHandler{
		serveMux: mux,
	}

	return http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: handler,
	}
}
