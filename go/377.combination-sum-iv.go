/*
 * @lc app=leetcode id=377 lang=golang
 *
 * [377] Combination Sum IV
 */

// @lc code=start
func combinationSum4(nums []int, target int) int {
	c := target
	dp := make([]int, c+1)
	dp[0] = 1
	for i := 1; i <= c; i++ {
		for _, num := range nums {
			if num <= i {
				dp[i] += dp[i - num]
			}
		}
	}
	return dp[c]
}
// @lc code=end

