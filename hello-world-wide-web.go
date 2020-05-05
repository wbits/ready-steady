package main

import (
	"fmt"
	"net/http"
)

func index(writer http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintf(writer, "Hello World, %s", request.URL.Path[1:])
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/foo", index)

	err := http.ListenAndServe(":3333", mux)
	if err != nil {
		fmt.Println(err)
	}
}
