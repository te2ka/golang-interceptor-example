package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world\n")
	})

	server.ListenAndServe()
}
