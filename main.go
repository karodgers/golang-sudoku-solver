package main

import (
	"fmt"
	"os"

	"sudoku/handlers"
)

const gridSize = 9

func main() {
	args := os.Args[1:]

	if len(args) != gridSize {
		fmt.Println("Error")
		return
	}

	board := make([][]int, gridSize)
	for i := range board {
		board[i] = make([]int, gridSize)
	}

	for i, arg := range args {
		if len(arg) != gridSize {
			fmt.Println("Error")
			return
		}

		for j, ch := range arg {
			if ch == '.' {
				board[i][j] = 0
			} else {
				num := int(ch - '0')
				if num < 1 || num > 9 {
					fmt.Println("Error")
					return
				}
				board[i][j] = num
			}
		}
	}

	if handlers.SolveSudoku(board) {
		handlers.PrintSudoku(board)
	} else {
		fmt.Println("Error")
	}
}
