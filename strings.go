package main

import "fmt"

func main() {
	fmt.Println(`Something in backticks
can spam multiple
lines`)
	fmt.Println("Something in double quotes is a string literal")

	var pi rune = 960
	smileyFace := rune(128515)
	a := rune(65)
	var vl byte = 255
	bang := byte(33)
	fmt.Printf("%v  %T  %c\n", pi, pi, pi)
	fmt.Printf("%v - %T\n", vl, vl)
	fmt.Printf("%c - %c - %c\n", smileyFace, a, bang)

	fmt.Printf("Byte alias for uint8 refers to an item from the 128-character subset of Unicode: ASCII")
	fmt.Printf("Byte alias for int32 refers an item from the 128-character subset of Unicode: ASCII")

	for c := byte(0); c <= 128; c++ {
		fmt.Printf("%3v: %U - %c\n", c, c, c)
	}

	a2 := rune('A')
	a3 := byte('A')
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Printf("%T - %T\n", a2, a3)

	peace := "vrede"
	peace = "shalom"
	fmt.Println(peace)
	fmt.Println(peace[5])
	fmt.Println(peace[2:])

	for i:=0; i<6; i++ {
		fmt.Printf("%v: %v -> %c\n", i, peace[i], peace[i])
	}

	s := "Julius Caesar"
	for i:=0; i<len(s); i++ {
		fmt.Printf("%c", s[i]+3)
	}
	fmt.Println()
	encoded := "Mxolxv#Fdhvdu"
	for i:=0; i<len(encoded); i++ {
		fmt.Printf("%c", encoded[i]-3)
	}
	fmt.Println()
}
