/*
 * @lc app=leetcode id=216 lang=golang
 *
 * [216] Combination Sum III
 */

// @lc code=start
func combinationSum3(k int, n int) [][]int {
	res := [][]int{}
	board := make([]bool, 9)
	var backtrack func(int, int, int)
	backtrack = func(idx, i, acc int) {
		if i == k {
			if acc == n {
				tmp := make([]int, k)
				t := 0
				for j := 0; j < 9; j++ {
					if board[j] {
						tmp[t] = j+1
						t++
					}
				}
				res = append(res, tmp)
			}
			return
		}
		for j := idx; j < 9; j++ {
			if acc+j+1 > n {
				break
			}
			board[j] = true
			backtrack(j+1, i+1, acc+j+1)
			board[j] = false
		}
	}

	backtrack(0, 0, 0)
	return res
}
// @lc code=end

