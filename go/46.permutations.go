/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start
func permute(nums []int) [][]int {
	res := [][]int{}
	if len(nums) == 0 {
		return res
	}
	bitmap := make([]bool, len(nums))
	var backtrack func([]int)
	backtrack = func(per []int) {
		if len(per) == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, per)
			res = append(res, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if bitmap[i] {
				continue
			}
			bitmap[i] = true
			per = append(per, nums[i])
			backtrack(per)
			per = per[:len(per)-1]
			bitmap[i] = false
		}
	}
	backtrack([]int{})
	return res
}
// @lc code=end

