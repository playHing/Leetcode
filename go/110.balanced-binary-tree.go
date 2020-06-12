/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
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
func isBalanced(root *TreeNode) bool {
	_, bal := dfsCheck(root)
	return bal
}

func dfsCheck(node *TreeNode) (height int, bal bool) {
	if node != nil {
		lh, lb := dfsCheck(node.Left)
		if !lb {
			return
		}
		rh, rb := dfsCheck(node.Right)
		if !rb || lh - rh < -1 || lh - rh > 1 {
			return
		}
		if lh < rh {
			height = 1 + rh
		} else {
			height = 1 + lh
		}
	}
	bal = true
	return
}

// @lc code=end

