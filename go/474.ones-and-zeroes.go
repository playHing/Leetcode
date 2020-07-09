/*
 * @lc app=leetcode id=474 lang=golang
 *
 * [474] Ones and Zeroes
 */

// @lc code=start
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	max := 0
	for _, str := range strs {
		n0, n1 := cnt01(str)
		for i := m; i >= n0; i-- {
			for j := n; j >= n1; j-- {
				if 1 + dp[i-n0][j-n1] > dp[i][j] {
					dp[i][j] = 1 + dp[i-n0][j-n1]
					if max < dp[i][j] {
						max = dp[i][j]
					}
				}
			}
		}
	}
	return max
}

func cnt01(str string) (m, n int) {
	for i := 0; i < len(str); i++ {
		if str[i] == '0' {
			m++
		} else {
			n++
		}
	}
	return
}
// @lc code=end
 
