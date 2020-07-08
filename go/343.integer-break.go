/*
 * @lc app=leetcode id=343 lang=golang
 *
 * [343] Integer Break
 */

// @lc code=start
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 1
	for i := 3; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			p1 := dp[j]
			if dp[j] < j {
				p1 = j
			}
			p2 := dp[i-j]
			if dp[i-j] < i-j {
				p2 = i-j
			}
			val := p1 * p2
			if dp[i] < val {
				dp[i] = val
			}
		}
	}
	return dp[n]
}
// @lc code=end

