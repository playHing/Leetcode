/*
 * @lc app=leetcode id=130 lang=golang
 *
 * [130] Surrounded Regions
 */

// @lc code=start
func solve(board [][]byte)  {
	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[0]))
	}

	var dfs func(int, int)
	dfs = func(i, j int) {
		invalid := i < 0 || i >= len(board)
		invalid = invalid || j < 0 || j >= len(board[0])
		invalid = invalid || board[i][j] == 'X' || visited[i][j]
		if invalid {
			return
		}
		visited[i][j] = true
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	for i := 0; i < len(board); i++ {
		if i == 0 || i == len(board) - 1 {
			for j := 0; j < len(board[0]); j++ {
				if board[i][j] == 'O' && !visited[i][j] {
					dfs(i, j)
				}
			}
		} else {
			for _, j := range []int{0, len(board[0])-1} {
				if board[i][j] == 'O' && !visited[i][j] {
					dfs(i, j)
				}
			}
		}
	}
	for i := 1; i < len(board)-1; i++ {
		for j := 1; j < len(board[0])-1; j++ {
			if board[i][j] == 'O' && !visited[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}
// @lc code=end

