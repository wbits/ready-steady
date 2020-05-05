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
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println(err)
	}
}
