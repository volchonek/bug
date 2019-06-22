package main

import (
	"fmt"
	"net/http"
)

func runServer(port string) {
	mux := http.NewServeMux()

	mux.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "port: ", port, "\nurl: ", r.URL.String())
		})

	mux.HandleFunc("/pong",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "pong")
		})

	server := http.Server{
		Addr:    port,
		Handler: mux,
	}

	fmt.Println("starting server at: ", port)

	server.ListenAndServe()
}
