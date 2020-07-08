/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 */

// @lc code=start
func rob(nums []int) int {
	max, notRobbedPrev := 0, 0
	for _, val := range nums {
		tmp := max
		if robCurrent := notRobbedPrev+val; robCurrent > max {
			max = robCurrent
		}
		notRobbedPrev = tmp
	}
	return max
}
// @lc code=end

