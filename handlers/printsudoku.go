package handlers

import "github.com/01-edu/z01"

func PrintSudoku(board [][]int) {
	for _, row := range board {
		for _, num := range row {
			z01.PrintRune(rune(num + '0'))
			z01.PrintRune(' ')
		}
		z01.PrintRune('\n')
	}
}
