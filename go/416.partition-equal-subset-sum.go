/*
 * @lc app=leetcode id=416 lang=golang
 *
 * [416] Partition Equal Subset Sum
 */

// @lc code=start
func canPartition(nums []int) bool {
	c := 0
	for _, v := range nums {
		c += v
	}
	if c % 2 == 1 {
		return false
	}
	c = c/2
	dp := make([]bool, c+1)
	dp[0] = true
	for i := 0; i < len(nums); i++ {
		for j := c; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}
	return dp[c]
}
// @lc code=end

