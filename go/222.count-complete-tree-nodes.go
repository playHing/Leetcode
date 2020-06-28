/*
 * @lc app=leetcode id=222 lang=golang
 *
 * [222] Count Complete Tree Nodes
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
func countNodes(root *TreeNode) int {
	left, right := root, root
	lh, rh := 0, 0
	for left != nil {
		left = left.Left
		lh++
	}
	for right != nil {
		right = right.Right
		rh++
	}
	if lh == rh {
		return (1 << lh) - 1
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
// @lc code=end

