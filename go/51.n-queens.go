/*
 * @lc app=leetcode id=51 lang=golang
 *
 * [51] N-Queens
 */

// @lc code=start
func solveNQueens(n int) [][]string {
	res := [][]string{}
	cols := make([]bool, n)
	dia1 := make([]bool, 2*n-1)
	dia2 := make([]bool, 2*n-1)
	rowtmp := []byte(strings.Repeat(".", n))
	var backtrack func(int, []string)
	backtrack = func(i int, board []string) {
		if i == n {
			tmp := make([]string, len(board))
			copy(tmp, board)
			res = append(res, tmp)
			return
		}
		for j := 0; j < n; j++ {
			if !cols[j] && !dia1[i+j] && !dia2[n-1+(j-i)] {
				cols[j], dia1[i+j], dia2[n-1+(j-i)] = true, true, true
				rowtmp[j] = 'Q'
				board = append(board, string(rowtmp))
				rowtmp[j] = '.'
				backtrack(i+1, board)
				board = board[:len(board)-1]
				cols[j], dia1[i+j], dia2[n-1+(j-i)] = false, false, false
			}
		}
	}

	backtrack(0, []string{})
	return res
}
// @lc code=end

