/*
 * @lc app=leetcode id=137 lang=golang
 *
 * [137] Single Number II
 */

// @lc code=start
func singleNumber(nums []int) int {
	x1, x2, mark := 0, 0, 0
	for _, v := range nums {
		x2 ^= x1 & v
		x1 ^= v
		mark = ^(x2 & x1) // k = 3 = '11'
		x2 &= mark
		x1 &= mark
	}
	return x1
}
// @lc code=end

