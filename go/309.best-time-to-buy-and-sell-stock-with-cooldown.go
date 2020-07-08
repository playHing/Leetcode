/*
 * @lc app=leetcode id=309 lang=golang
 *
 * [309] Best Time to Buy and Sell Stock with Cooldown
 */

// @lc code=start
func maxProfit(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	dp := make([]int, n)
	for i := n-2; i >= 0; i-- {
		dp[i] = dp[i+1]

		for j := i+1; j < n; j++ {
			val := prices[j]-prices[i]
			if j+2 < n {
				val += dp[j+2]
			}
			if dp[i] < val {
				dp[i] = val
			}
		}
	}
	return dp[0]
}
// @lc code=end

