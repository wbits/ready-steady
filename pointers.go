package main

import (
	"fmt"
)

type something *string
func foo(s something) somthing {
	return s
}


func main() {
	answer := 42
	memoryAddress := &answer

	fmt.Println(*memoryAddress == answer)



}
