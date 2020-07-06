/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */

// @lc code=start
func subsets(nums []int) [][]int {
	res := [][]int{}
	var backtrack func(int, []int)
	backtrack = func(idx int, comb []int) {
		tmp := make([]int, len(comb))
		copy(tmp, comb)
		res = append(res, tmp)
		
		for i := idx; i < len(nums); i++ {
			comb = append(comb, nums[i])
			backtrack(i+1, comb)
			comb = comb[:len(comb)-1]
		}
	}

	backtrack(0, []int{})
	return res
}
// @lc code=end

