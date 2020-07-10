/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	max, n := 1, len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] >= nums[i] {
				continue
			}
			if dp[i] < dp[j] + 1 {
				dp[i] = dp[j] + 1
				if max < dp[i] {
					max = dp[i]
				}
			}
		}
	}
	return max
}
// @lc code=end

