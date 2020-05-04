package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 80
	height = 15
)

type Universe [][]bool

func main() {
	u := NewUniverse()
	u.Seed()
	u.Show()

	steps := 0

	for {
		steps++
		time.Sleep(time.Second)
		Step(u, NewUniverse())
		u.Show()
		if steps == 30 {
			break
		}
	}


}

func NewUniverse() Universe {
	u := make(Universe, height)
	for r := range u {
		u[r] = make([]bool, width)
	}

	return u
}

func (u Universe) Seed() {
	for row := range u {
		for cell := range u[row] {
			rand.Seed(time.Now().UnixNano())
			u[row][cell] = rand.Intn(4) == 0
		}
	}
}

func (u Universe) Show() {
	print("\x0c")
	for r := 0; r < height; r++ {
		row := ""
		for c := 0; c < width; c++ {
			if u[r][c] {
				row += " * "
			} else {
				row += "   "
			}
		}
		fmt.Println(row)
	}
}


func (u Universe) Alive(x, y int) bool {
	b := (x+width) % width
	a := (y+height) % height

	return u[a][b]
}

func (u Universe) Neighbors(x, y int) int {
	count := 0
	for i := -1; i<2; i++ {
		c := i + x
		r := i + y
		if c == x && r == y {
			continue
		}
		if u.Alive(c, r) {
			count++
		}
	}

	return count
}

func (u Universe) Next(x, y int) bool {
	// A live cell with less than two live neighbors dies.
	//A live cell with two or three live neighbors lives on to the next generation.
	//A live cell with more than three live neighbors dies.
	//A dead cell with exactly three live neighbors becomes a live cell.
	neighbors := u.Neighbors(x, y)
	if u.Alive(x, y) {
		if neighbors < 2 || neighbors > 3 {
			return false
		}
		return true
	}

	return neighbors == 3
}


func Step(a, b Universe) {
	for r := range a {
		for c := range a[r] {
			b[r][c] = a.Next(c, r)
		}
	}

	a = b
	a.Show()
}
