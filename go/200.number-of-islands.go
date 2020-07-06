/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	num := 0
	visited := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}
	var dfs func(int, int)
	dfs = func(i, j int) {
		invalid := i < 0 || i >= len(grid)
		invalid = invalid || j < 0 || j >= len(grid[0])
		invalid = invalid || grid[i][j] == '0' || visited[i][j] == true
		if invalid {
			return
		}
		visited[i][j] = true
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if !visited[i][j] && grid[i][j] == '1' {
				dfs(i, j)
				num++
			}
		}
	}
	return num
}
// @lc code=end