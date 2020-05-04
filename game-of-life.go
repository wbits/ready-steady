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

type Universe []Cell
type Cell struct {
	row int
	col int
	alive bool
}

func main() {
	a, b := NewUniverse(), NewUniverse()
	a.Seed()

	for i := 0; i < 30; i++ {
		a.Show()
		Step(a, b)
		time.Sleep(time.Second / 2)
		a, b = b, a // Swap universes
	}
}

func NewUniverse() Universe {
	u := make(Universe, height*width)
	for i := range u {
		u[i] = NewCell(i%height, i%width)
	}

	return u
}

func NewCell(row int, col int) Cell {
	c := Cell{}
	c.row = row
	c.col = col
	c.alive = false
	return c
}

func (u Universe) Seed() {
	for row := range u {
		rand.Seed(time.Now().UnixNano())
		u[row].alive = rand.Intn(4) == 0
	}
}

func (u Universe) String() string {
	buffer := make([]byte, 0, len(u))
	for index, cell := range u {
		byteCharacter := byte(' ')
		if cell.alive {
			byteCharacter = '*'
		}
		buffer = append(buffer, byteCharacter)
		if index%width == 0 {
			buffer = append(buffer, '\n')
		}
	}

	return string(buffer)
}

// Show clears the screen and displays the universe.
func (u Universe) Show() {
	fmt.Print("\x0c", u.String())
}


func (u Universe) Alive(column, row int) bool {
	return u.Cell(column, row).alive
}

func (u Universe) Cell(column, row int) Cell {
	return u[u.Index(column, row)]
}

func (u Universe) Index(column, row int) int {
	column = (column + width) % width
	row = (row + height) % height

	return (row * width) + column
}

func (u Universe) Neighbors(column, row int) int {
	count := 0
	for i := -1; i<=1; i++ {
		c := i + column
		r := i + row
		cell := u.Cell(c, r)
		if !(c == column && r == row) && cell.alive {
			count++
		}
	}

	return count
}

func (u Universe) Next(column, row int) bool {
	neighbors := u.Neighbors(column, row)
	if u.Alive(column, row) {
		if neighbors < 2 || neighbors > 3 {
			return false
		}
		return true
	}

	return neighbors == 3
}


func Step(a, b Universe) {
	for i, cell := range a {
		b[i].alive = a.Next(cell.col, cell.row)
	}
}
