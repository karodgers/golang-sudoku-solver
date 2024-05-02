package handlers

func SolveSudoku(board [][]int) bool {
	empty := FindEmptyCell(board)
	if empty == [2]int{-1, -1} {
		return true // Sudoku solved
	}

	row, col := empty[0], empty[1]

	for num := 1; num <= 9; num++ {
		if IsValid(board, row, col, num) {
			board[row][col] = num

			if SolveSudoku(board) {
				return true
			}

			board[row][col] = 0
		}
	}

	return false
}
