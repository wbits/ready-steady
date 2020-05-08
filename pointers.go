package main

import (
	"fmt"
)

type person struct {
	name, superpower string
	age              int
}

func somePointer() *string {
	v := "foo"

	return &v
}

func main() {
	answer := 42
	memoryAddress := &answer
	fmt.Println(*memoryAddress == answer)

	var administrator *string
	scolese := "Christopheer J. Scolese"
	administrator = &scolese
	fmt.Println(*administrator, administrator)

	bolden := "Charles F. Bolden"
	administrator = &bolden
	fmt.Println(*administrator, administrator)

	bolden = "foo"
	scolese = "bar"
	fmt.Println(*administrator, administrator)
	administrator = &scolese
	fmt.Println(*administrator, administrator)

	timmy := &person{name: "Timothy", age: 10}
	timmy.superpower = "flying"
	fmt.Printf("%+v\n", timmy)

	superpowers := &[3]string{"flying", "invisibility", "super strength"}
	fmt.Println(superpowers[0])
	fmt.Println(superpowers[2:3])
	fmt.Println(*somePointer())
}
