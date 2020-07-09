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
	m := 2*total + 1
	
	if j := S + total; j >= m {
	    return 0
	}
	
	dp := make([][]int, 2)
	for i := range dp {
	    dp[i] = make([]int, m)
	}
	dp[0][0+total] = 1
	for i := 0; i < n; i++ {
	    for j := 0; j < m ; j++ {
		    if dp[0][j] > 0 {
			    for sign := -1; sign <= 1; sign += 2 {
					dp[1][j + sign*nums[i]] += dp[0][j]
				}
			}
		}
		dp[0], dp[1] = dp[1], dp[0] // swap the two rows
		for j := 0; j < m; j++ {
		    dp[1][j] = 0 // clear the new second row for the next iteration
		}
	}
	return dp[0][S+total]
}
// @lc code=end

