/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := [][]int{}
	var backtrack func(int, int, []int)
	backtrack = func(idx, acc int, comb []int) {
		if acc == target {
			tmp := make([]int, len(comb))
			copy(tmp, comb)
			res = append(res, tmp)
			return
		}
		for i := idx; i < len(candidates); {
			cur := candidates[i]
			if acc + cur > target {
				break
			}
			ni := i+1
			for ni < len(candidates) && candidates[ni] == cur {
				ni++
			}
			originsize, acc2 := len(comb), acc
			for j := i; j < ni; j++ {
				acc2 += cur
				if acc2 > target {
					break
				}
				comb = append(comb, cur)
				backtrack(ni, acc2, comb)
			}
			comb = comb[:originsize]
			i = ni
		}
	}
	
	backtrack(0, 0, []int{})
	return res
}
// @lc code=end

