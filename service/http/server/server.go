package server

import (
	"fmt"
	"net/http"

	"github.com/te2ka/golang-interceptor-example/service/http/interceptor"
)

func Setup() {
	http.HandleFunc("/", interceptor.ApplyInterceptors(helloHandler, interceptor.WithCountConnection(), interceptor.WithLogging()))
}

func helloHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Hello, world")
}
