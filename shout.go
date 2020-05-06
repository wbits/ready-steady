package main

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}

type laser struct {
}

type spaceship struct {
	laser
}

func shout(t talker) string {
	return strings.ToUpper(t.talk())
}

func (l laser) talk() string {
	return strings.Repeat("pew", 4)
}

func main() {
	laser := laser{}
	fmt.Println(shout(laser))

	spaceship := spaceship{laser}
	fmt.Println(shout(spaceship))
}
