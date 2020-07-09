/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	c := amount
	dp := make([]int, c+1)
	if c <= 0 {
		return 0
	}
	if len(coins) == 0 {
		return -1
	}

	// Normal DP
	// for j := 1; j <= c; j++ {
	// 	dp[j] = c+1
	// 	for i := 0; i < len(coins); i++ {
	// 		for num := 1; num * coins[i] <= j; num++ {
	// 			if dp[j - num * coins[i]] == c+1 {
	// 				continue
	// 			}
	// 			if dp[j - num * coins[i]] + num < dp[j] {
	// 				dp[j] = dp[j - num * coins[i]] + num
	// 			}
	// 		}
	// 	}
	// }
	// if dp[c] == c+1 {
	// 	return -1
	// }
	// return dp[c]

	// 0-1 package DP
	for i := 1; i <= c; i++ {
		dp[i] = c+1
	}
	for i := 0; i < len(coins); i++ {
		for num := 1; num * coins[i] <= c; num++ {
			for j := c; j >= num * coins[i]; j-- {
				if dp[j - num * coins[i]] + num < dp[j] {
					dp[j] = dp[j - num * coins[i]] + num
				}
			}
		}
	}
	if dp[c] == c+1 {
		return -1
	}
	return dp[c]
}
// @lc code=end

