/*
 * @lc app=leetcode id=494 lang=golang
 *
 * [494] Target Sum
 */

// @lc code=start
func findTargetSumWays(nums []int, S int) int {
	n := len(nums)
	total := 0
	for _, v := range nums {
		total += v
	}
	if total < S || S < -total {
		return 0
	}
	c := 2 * total + 1
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, c)
	}
	dp[0][total] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < c; j++ {
			if dp[0][j] == 0 {
				continue
			}
			for sign := -1; sign <= 1; sign += 2 {
				dp[1][j+sign*nums[i]] += dp[0][j]
			}
		}

		dp[0], dp[1] = dp[1], dp[0]
		for k := range dp[1] {
			dp[1][k] = 0
		}
	}
	return dp[0][total+S]
}
// @lc code=end

