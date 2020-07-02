/*
 * @lc app=leetcode id=230 lang=golang
 *
 * [230] Kth Smallest Element in a BST
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
func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	stack := []*TreeNode{}
	for root != nil {
		stack = append(stack, root)
		root = root.Left
	}

	for k != 1 {
		node := stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]

		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		k = k-1
	}

	return stack[len(stack)-1].Val
}
// @lc code=end

