package main

import "fmt"

type number struct {
	value int
	valid bool
}

func newNumber(n int) number {
	return number{n, true}
}

func (n number) String() string {
	if !n.valid {
		return "not set"
	}

	return fmt.Sprintf("%d", n.value)
}

func main() {
	n := newNumber(12)
	fmt.Println(n)

	e := number{}
	fmt.Println(e)
}
