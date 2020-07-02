/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := [][]int{}
	var backtrack func(int, int, []int)
	backtrack = func(start, acc int, comb []int) {
		if acc == target {
			tmp := make([]int, len(comb))
			copy(tmp, comb)
			res = append(res, tmp)
			return
		}
		for i := start; i < len(candidates) && candidates[i] + acc <= target; i++ {
			cnt, acc2 := 1, candidates[i] + acc
			for acc2 <= target {
				comb = append(comb, candidates[i])
				backtrack(i+1, acc2, comb)
				cnt, acc2 = cnt+1, candidates[i] + acc2
			}
			comb = comb[:len(comb)-cnt+1]
		}
	}
	backtrack(0, 0, []int{})
	return res
}
// @lc code=end

