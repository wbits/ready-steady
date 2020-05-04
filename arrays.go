package main

import (
	"fmt"
)

type planet string
type galaxy []planet

func main() {
	g := galaxy([]planet{"Mercury", "Venus", "Earth", "Jupiter",
		"Mars", "Saturn", "Uranus", "Neptune",
	})

	fmt.Println(g.terraform("Yo Rock "))
}

func (g galaxy) terraform(prefix string) galaxy {
	newGalaxy := make(galaxy, len(g))
	for i, p := range g {
		newGalaxy[i] = p.prefix(prefix)
	}

	return newGalaxy
}

func (p planet) prefix(prefix string) planet {
	return planet(prefix + string(p))
}
