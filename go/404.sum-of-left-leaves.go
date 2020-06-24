/*
 * @lc app=leetcode id=404 lang=golang
 *
 * [404] Sum of Left Leaves
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
func sumOfLeftLeaves(root *TreeNode) int {
	return recal(root, false)
}

func recal(root *TreeNode, isLeft bool) int {
	if root == nil {
		return 0
	}
	if isLeft && root.Left == nil && root.Right == nil {
		return root.Val
	}
	return recal(root.Left, true) + recal(root.Right, false)
}
// @lc code=end

