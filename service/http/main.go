package main

import (
	"net/http"

	"github.com/te2ka/golang-interceptor-example/service/http/server"
)

func main() {
	s := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	server.Setup()

	s.ListenAndServe()
}
