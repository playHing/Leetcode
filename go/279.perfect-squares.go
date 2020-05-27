/*
 * @lc app=leetcode id=279 lang=golang
 *
 * [279] Perfect Squares
 */

// @lc code=start
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			if i-j*j == 0 {
				dp[i] = 1
			} else {
				if dp[i] == 0 {
					dp[i] = dp[i-j*j] + 1
				} else if dp[i-j*j]+1 < dp[i] {
					dp[i] = dp[i-j*j] + 1
				}
			}
		}
	}
	return dp[n]
}

// @lc code=end

