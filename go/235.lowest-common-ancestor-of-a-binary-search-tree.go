/*
 * @lc app=leetcode id=235 lang=golang
 *
 * [235] Lowest Common Ancestor of a Binary Search Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val > q.Val {
		p, q = q, p
	}
	for root != nil {
		if p.Val <= root.Val && root.Val <= q.Val {
			return root
		}
		if p.Val <= root.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return nil
}
// @lc code=end

