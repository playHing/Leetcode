/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 */

// @lc code=start
func combine(n int, k int) [][]int {
	res := [][]int{}
	var backtrack func(int, []int)
	backtrack = func(idx int, comb []int) {
		if len(comb) == k {
			tmp := make([]int, k)
			copy(tmp, comb)
			res = append(res, tmp)
			return
		}

		if len(comb)+(n-idx+1) < k {
			return
		}
		
		for i := idx; i <= n; i++ {
			comb = append(comb, i)
			backtrack(i+1, comb)
			comb = comb[:len(comb)-1]
		}
	}

	backtrack(1, []int{})
	return res
}
// @lc code=end

