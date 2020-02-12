/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 */

// @lc code=start
func sortColors(nums []int) {
	l, r := 0, len(nums)-1
	for i := 0; i <= r && l < r; {
		switch nums[i] {
		case 0:
			nums[i] = nums[l]
			nums[l] = 0
			l++
			if i < l {
				i = l
			}
		case 2:
			nums[i] = nums[r]
			nums[r] = 2
			r--
		default:
			i++
		}
	}
}

// @lc code=end

