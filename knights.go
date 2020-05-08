package main

import "fmt"

type item struct {
	name string
}

type character struct {
	name     string
	leftHand *item
}

func (c *character) pickup(i *item) {
	c.leftHand = i
	fmt.Printf("%v picks up a %v\n", c.name, i)
}

func (c *character) give(to *character) {
	to.leftHand = c.leftHand
	fmt.Printf("%v gives a %v to %v\n", c.name, c.leftHand, to.name)
}

func main() {
	something := &item{"something"}
	arthur := &character{name: "Arthur"}
	knight := &character{name: "Knight"}

	arthur.pickup(something)
	arthur.give(knight)

	fmt.Println(arthur.leftHand)
	fmt.Println(knight.leftHand)
}
