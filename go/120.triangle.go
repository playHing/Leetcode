/*
 * @lc app=leetcode id=120 lang=golang
 *
 * [120] Triangle
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {
	for i := 1; i < len(triangle); i++ {
		triangle[i][0] += triangle[i-1][0]
		for j := 1; j < len(triangle[i])-1; j++ {
			sum := triangle[i-1][j-1]
			if triangle[i-1][j] < sum {
				sum = triangle[i-1][j]
			}
			triangle[i][j] += sum
		}
		triangle[i][len(triangle[i])-1] += triangle[i-1][len(triangle[i-1])-1]
	}
	min := triangle[len(triangle)-1][0]
	for i := 1; i < len(triangle[len(triangle)-1]); i++ {
		if triangle[len(triangle)-1][i] < min {
			min = triangle[len(triangle)-1][i]
		}
	}
	return min
}
// @lc code=end

