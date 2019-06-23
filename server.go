package main

import (
	"fmt"
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong") // заменить на реальный ответ от btest
}

func runServer(port string) {
	mux := http.NewServeMux()

	mux.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "port: ", port, "\nurl: ", r.URL.String())
		})

	mux.HandleFunc("/ping", ping)

	server := http.Server{
		Addr:    port,
		Handler: mux,
	}

	fmt.Println("starting server at ", port)

	server.ListenAndServe()
}
