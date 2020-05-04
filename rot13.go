package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	message := "uv vagreangvbany fcnpr fgngvba"
	for _, c := range message {
		if c >= 'a' && c <= 'z' {
			c += 13
		}
		if c > 'z' {
			c = c-26
		}
		fmt.Printf("%c", c)
	}

	question := "¿Cómo estás?"

	fmt.Println(len(question), "bytes")
	fmt.Println(utf8.RuneCountInString(question), "runes")

	c, size := utf8.DecodeRuneInString(question)
	fmt.Printf("First rune: %c %v bytes", c, size)

	for _, c := range question {
		fmt.Printf("%v %T %c\n", c, c, c)
	}

	fmt.Println(utf8.RuneCountInString("abcdefghijklmnopqrstuvwxyz"))
}
