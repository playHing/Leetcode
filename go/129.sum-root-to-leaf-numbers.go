/*
 * @lc app=leetcode id=129 lang=golang
 *
 * [129] Sum Root to Leaf Numbers
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
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
    return dfs(root, 0)
}

func dfs(root *TreeNode, acc int) int {
	acc = 10 * acc + root.Val
	if root.Left == nil && root.Right == nil {
		return acc
	}
	acc2 := 0
	if root.Left != nil {
		acc2 += dfs(root.Left, acc)
	}
	if root.Right != nil {
		acc2 += dfs(root.Right, acc)
	}
	return acc2
}
// @lc code=end

