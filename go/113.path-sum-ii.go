/*
 * @lc app=leetcode id=113 lang=golang
 *
 * [113] Path Sum II
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
    if root == nil {
		return [][]int{}
	}
	return pathSumStep(root, sum, []int{})
}

func pathSumStep(root *TreeNode, sum int, path []int) [][]int {
	sum -= root.Val
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil {
		if sum == 0 {
			path2 := make([]int, len(path))
			copy(path2, path)
			return [][]int{path2}
		}
		return [][]int{}
	}
	res := [][]int{}
	if root.Left != nil {
		res = append(res, pathSumStep(root.Left, sum, path)...)
	}
	if root.Right != nil {
		res = append(res, pathSumStep(root.Right, sum, path)...)
	}
	return res
}
// @lc code=end

