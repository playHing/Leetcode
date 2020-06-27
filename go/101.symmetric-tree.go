/*
 * @lc app=leetcode id=101 lang=golang
 *
 * [101] Symmetric Tree
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
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
    return isMirror(root.Left, root.Right)
}

func isMirror(r1, r2 *TreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	}
	if r1 == nil || r2 == nil || r1.Val != r2.Val {
		return false
	}
	return isMirror(r1.Left, r2.Right) && isMirror(r1.Right, r2.Left)
}
// @lc code=end

