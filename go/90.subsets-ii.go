/*
 * @lc app=leetcode id=90 lang=golang
 *
 * [90] Subsets II
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	var backtrack func(int, []int)
	backtrack = func(idx int, comb []int) {
		res = append(res, append([]int{}, comb...))

		for i := idx; i < len(nums); i++ {
			comb = append(comb, nums[i])
			backtrack(i+1, comb)
			comb = comb[:len(comb)-1]
			for i+1 < len(nums) && nums[i+1] == nums[i] {
				i++
			}
		}
	}

	backtrack(0, []int{})
	return res
}
// @lc code=end

