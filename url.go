package main

import (
	"fmt"
	"net/url"
)

func main() {
	_, err := url.Parse("https://a b.com/")

	if err != nil {
		fmt.Printf("%#v - %v\n", err, err.Error())

		if e, ok := err.(*url.Error); ok {
			fmt.Println("Op ", e.Op)
			fmt.Println("Err ", e.Err)
			fmt.Println("Url: ", e.URL)
		}
	}
}
