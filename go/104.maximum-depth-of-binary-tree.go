/*
 * @lc app=leetcode id=104 lang=golang
 *
 * [104] Maximum Depth of Binary Tree
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
func maxDepth(root *TreeNode) int {
    if root == nil {
		return 0
	}
	ld, rd := maxDepth(root.Left), maxDepth(root.Right)
	if ld < rd {
		return rd + 1
	} else {
		return ld + 1
	}
}
// @lc code=end

