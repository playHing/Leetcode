/*
 * @lc app=leetcode id=145 lang=golang
 *
 * [145] Binary Tree Postorder Traversal
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
func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
    if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		node := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if node.Left == nil && node.Right == nil {
			res = append(res, node.Val)
			continue
		}
		queue = append(queue, node)
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		node.Left, node.Right = nil, nil
	}
	return res
}
// @lc code=end

