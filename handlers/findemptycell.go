package handlers

func FindEmptyCell(board [][]int) [2]int {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{-1, -1} // No empty cell found
}
