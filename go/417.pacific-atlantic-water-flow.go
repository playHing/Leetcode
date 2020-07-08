/*
 * @lc app=leetcode id=417 lang=golang
 *
 * [417] Pacific Atlantic Water Flow
 */

// @lc code=start


var neigborDist = [][]int{[]int{-1,0},[]int{1,0},[]int{0,-1},[]int{0,1}}

func pacificAtlantic(matrix [][]int) [][]int {
	res := [][]int{}
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return res
	}
	toPacific := make([][]bool, len(matrix))
	toAtlantic := make([][]bool, len(matrix))
	for i := 0; i < len(matrix); i++ {
		toPacific[i] = make([]bool, len(matrix[0]))
		toAtlantic[i] = make([]bool, len(matrix[0]))
	}
	for i := 0; i < len(matrix); i++ {
		dfs(i, 0, matrix[i][0], &matrix, &toPacific)
		dfs(i, len(matrix[0])-1, matrix[i][len(matrix[0])-1], &matrix, &toAtlantic)
	}
	for j := 0; j < len(matrix[0]); j++ {
		dfs(0, j, matrix[0][j], &matrix, &toPacific)
		dfs(len(matrix)-1, j, matrix[len(matrix)-1][j], &matrix, &toAtlantic)
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if toPacific[i][j] && toAtlantic[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}

func dfs(i, j, prev int, mtx *[][]int, toboundary *[][]bool) {
	matrix := *mtx
	if i < 0 || i >= len(matrix) {
		return
	}
	if j < 0 || j >= len(matrix[0]) {
		return
	}
	if matrix[i][j] < prev || (*toboundary)[i][j] {
		return
	}
	(*toboundary)[i][j] = true
	dfs(i-1, j, matrix[i][j], mtx, toboundary)
	dfs(i+1, j, matrix[i][j], mtx, toboundary)
	dfs(i, j-1, matrix[i][j], mtx, toboundary)
	dfs(i, j+1, matrix[i][j], mtx, toboundary)
}