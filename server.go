package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	url := "http://192.168.11.80:8082/ping?param=ping"
	resp, err := http.Get(url)
	if err != nil {
		//printError(err)
		fmt.Println(fmt.Sprint(err))
		return
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(respBody))
		return
	}

	fmt.Printf("http.Get body %#v\n\n\n", string(respBody))
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
