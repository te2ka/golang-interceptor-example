package main

import (
	"github.com/te2ka/golang-interceptor-example/service/http/server"
	_ "github.com/te2ka/golang-interceptor-example/utils/log"
)

func main() {
	s := server.NewServer(8080)

	s.ListenAndServe()
}
