/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	var dfs func(int, int, int) bool
	dfs = func(i, j, k int) bool {
		invalid := i < 0 || i >= len(board)
		invalid = invalid || j < 0 || j >= len(board[0])
		if invalid {
			return false
		}
		if k == len(word)-1 {
			return word[k] == board[i][j]
		}
		if word[k] != board[i][j] {
			return false
		}
		tmp := board[i][j]
		board[i][j] = '#'
		res := dfs(i-1, j, k+1) || dfs(i+1, j, k+1) || dfs(i, j-1, k+1) || dfs(i, j+1, k+1)
		board[i][j] = tmp
		return res
	}
    for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
// @lc code=end

