/*
 * @lc app=leetcode id=209 lang=golang
 *
 * [209] Minimum Size Subarray Sum
 */

// @lc code=start
func minSubArrayLen(s int, nums []int) int {
	min := len(nums) + 1
	sum := 0
	for i, j := 0, 0; i < len(nums); {
		if sum < s && j < len(nums) {
			sum += nums[j]
			j++
		} else {
			sum -= nums[i]
			i++
		}
		if sum >= s && min > j-i {
			min = j - i
		}
	}
	if min == len(nums)+1 {
		return 0
	} else {
		return min
	}
}

// @lc code=end

