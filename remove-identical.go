package main

import "fmt"

func source(downstream chan string) {
	list := []string{"one", "two", "two", "three", "three", "three", "four"}
	for _, item := range list {
		downstream <- item
	}

	close(downstream)
}

func removeDuplicate(upstream, downstream chan string) {
	var previousItem string
	for item := range upstream {
		if item == previousItem {
			continue
		}

		previousItem = item
		downstream <- item
	}

	close(downstream)
}

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go source(channel1)
	go removeDuplicate(channel1, channel2)

	for item := range channel2 {
		fmt.Println(item)
	}
}
