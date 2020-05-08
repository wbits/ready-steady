package main

import "fmt"

type turtle struct {
	x, y int
}

func (p *turtle) up() {
	p.y++
}

func (p *turtle) down() {
	p.y--
}

func (p *turtle) right() {
	p.x++
}

func (p *turtle) left() {
	p.x--
}

func main() {
	var t turtle
	t.up()
	t.right()
	t.up()

	fmt.Printf("%+v", t)
}
