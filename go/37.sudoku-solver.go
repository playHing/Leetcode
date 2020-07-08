/*
 * @lc app=leetcode id=37 lang=golang
 *
 * [37] Sudoku Solver
 */

// @lc code=start
func solveSudoku(board [][]byte)  {
	
	isValid := func(i, j int, val byte) bool {
		for k := 0; k < 9; k++ {
			if board[i][k] == val {
				return false
			}
			if board[k][j] == val {
				return false
			}
			if board[(i/3)*3+k/3][(j/3)*3+k%3] == val {
				return false
			}
		}
		return true
	}

	var solve func() bool
	solve = func() bool {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if board[i][j] != '.' {
					continue
				}
				for k := 1; k <= 9; k++ {
					if isValid(i, j, byte(k)+'0') {
						board[i][j] = byte(k)+'0'
						if solve() {
							return true
						}
						board[i][j] = '.'
					}
				}
				return false
			}
		}
		return true
	}

	solve()
}
// @lc code=end

