/*
 * @lc app=leetcode id=64 lang=golang
 *
 * [64] Minimum Path Sum
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[0]))
		dp[i][0] = grid[i][0]
		if i > 0 {
			dp[i][0] += dp[i-1][0]
		}
	}
	for j := range dp[0] {
		dp[0][j] = grid[0][j]
		if j > 0 {
			dp[0][j] += dp[0][j-1]
		}
	}
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			if dp[i-1][j] < dp[i][j-1] {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			}
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}
// @lc code=end

