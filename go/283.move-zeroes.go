/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 */

// @lc code=start
func moveZeroes(nums []int)  {
	if nums == nil || len(nums) == 0 {
		return	
	}
	boxIdx := 0
	for ballIdx := 0; ballIdx < len(nums); ballIdx++ {
		if nums[ballIdx] != 0 {
			if boxIdx < ballIdx {
				nums[boxIdx] = nums[ballIdx]
				nums[ballIdx] = 0
			}
			boxIdx++
		}
	}
}
// @lc code=end

