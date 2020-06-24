/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
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
func isValidBST(root *TreeNode) bool {
    return validate(root, 0, 0, false, false)
}

func validate(root *TreeNode, lb, rb int, lv, rv bool) bool {
	if root == nil {
		return true
	}
	if (!lv || lb < root.Val) && (!rv || root.Val < rb) {
		return validate(root.Left, lb, root.Val, lv, true) && validate(root.Right, root.Val, rb, true, rv)
	}
	return false
}
// @lc code=end

