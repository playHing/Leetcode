/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	loc := make(map[int]int)
	for i, v := range nums {
		if j, b := loc[v]; b {
			return []int{j, i}
		} else {
			loc[target-v] = i
		}
	}
	return nil
}

// @lc code=end

