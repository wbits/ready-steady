package main

import (
	"errors"
	"fmt"
	"strings"
)

const rows, columns = 9, 9

var (
	ErrOutOfBounds = errors.New("outside of grid boundaries")
	ErrInvalidDigit = errors.New("invalid digit given")
)

type (
	grid [rows][columns]int8
	SudokuError []error
)

func (e SudokuError) Error() string {
	var s []string
	for _, err := range e {
		s = append(s, err.Error())
	}
	return strings.Join(s, ", ")
}

func (g grid) set(row, column int, digit int8) error {
	var errs SudokuError
	if !inbound(row, column) {
		errs = append(errs, ErrOutOfBounds)
	}

	if !validDigit(digit) {
		errs = append(errs, ErrInvalidDigit)
	}

	if len(errs) > 0 {
		return errs
	}

	g[row][column] = digit

	return nil
}

func inbound(r, c int) bool {
	return r >= 0 && r < rows && c >= 0 && c < columns
}

func validDigit(d int8) bool {
	return d > 0 && d <= 9
}

func main() {
	g := grid{}
	err := g.set(8, 8, 0)
	if err != nil {
		fmt.Printf("An error occurred: %v.\n", err)
	}
}
