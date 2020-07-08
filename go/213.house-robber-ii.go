/*
 * @lc app=leetcode id=213 lang=golang
 *
 * [213] House Robber II
 */

// @lc code=start
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	
	max, prev := 0, 0
	for i := 2; i < len(nums)-1; i++ {
		tmp := max
		if max < prev + nums[i] {
			max = prev + nums[i]
		}
		prev = tmp
	}
	planA := max + nums[0]

	max, prev = 0, 0
	for i := 1; i < len(nums); i++ {
		tmp := max
		if max < prev + nums[i] {
			max = prev + nums[i]
		}
		prev = tmp
	}
	planB := max

	if planA > planB {
		return planA
	} else {
		return planB
	}
}
// @lc code=end

