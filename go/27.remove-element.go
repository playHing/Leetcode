/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	cnt, ptr := 0, 0
	for _, v := range nums {
		if v != val {
			nums[ptr] = v
			ptr++
			cnt++
		}
	}
	return cnt
}

// @lc code=end

