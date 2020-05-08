package main

import (
	"fmt"
	"strings"
)

func splitWords(sentence string, c chan string) {
	words := strings.Fields(sentence)
	for _, word := range words {
		c <- strings.Trim(word, ".?!;:,")
	}

	close(c)
}

func main() {
	c := make(chan string)
	go splitWords("The pig stood right up in the brook.", c)
	for word := range c {
		fmt.Println(word)
	}
}
