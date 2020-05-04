package main

import (
	"fmt"
	"strings"
)

type board [8][8]rune

func main() {
	drawBoard(initBoard())
}

func drawBoard(b board) {
	for _, row := range b {
		for _, piece := range row {
			if piece >= 'A' {
				fmt.Printf("%4c", piece)
			} else {
				fmt.Printf("%4c", ' ')
			}
		}
		fmt.Println()
	}
}

func initBoard() board {
	var b board
	blackPieces := "rnbkqbnr"
	whitePieces := strings.ToUpper(blackPieces)

	for i, piece := range blackPieces {
		b[0][i] = piece
		b[1][i] = 'p'
	}

	for i, piece := range whitePieces {
		b[6][i] = 'p'
		b[7][i] = piece
	}

	return b
}
