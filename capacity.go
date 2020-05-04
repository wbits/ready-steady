package main

import "fmt"

func main() {
	list := []string{"foo", "bar"}

	for i := rune(0); i<10; i++ {
		list = append(list, string(i + 'a'))
		fmt.Println(cap(list), list)
	}
}
