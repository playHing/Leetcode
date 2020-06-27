/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
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
func inorderTraversal(root *TreeNode) []int {
    if root == nil {
		return nil
	}
	res := make([]int, 0)
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		node := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if node.Left == nil && node.Right == nil {
			res = append(res, node.Val)
			continue
		}
		right, left := node.Right, node.Left
		node.Right, node.Left = nil, nil
		if right != nil {
			queue = append(queue, right)
		}
		queue = append(queue, node)
		if left != nil {
			queue = append(queue, left)
		}
	}
	return res
}
// @lc code=end

