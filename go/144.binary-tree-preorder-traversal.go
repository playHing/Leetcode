/*
 * @lc app=leetcode id=144 lang=golang
 *
 * [144] Binary Tree Preorder Traversal
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
func preorderTraversal(root *TreeNode) []int {
    if root == nil {
		return nil
	}
	res := []int{}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		node := queue[len(queue)-1]
		res = append(res, node.Val)
		queue = queue[:len(queue)-1]
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
	}
	return res
}
// @lc code=end

