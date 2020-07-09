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
	dp := make([]int, c+1)
	for i := 0; i < len(nums); i++ {
		for j := c; j >= 0; j-- {
			if nums[i] <= j && dp[j-nums[i]] + nums[i] > dp[j] {
				dp[j] = dp[j-nums[i]] + nums[i]
			}
		}
	}
	fmt.Println(dp[c], c)
	return dp[c] == c
}
// @lc code=end

