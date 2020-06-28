/*
 * @lc app=leetcode id=450 lang=golang
 *
 * [450] Delete Node in a BST
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
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {

		if root.Right == nil {
			return root.Left
		} else if root.Left == nil {
			return root.Right
		} else {

			prev, tmp := root, root.Right
			if tmp.Left == nil {
				root.Val = tmp.Val
				root.Right = tmp.Right
			} else {
				for tmp.Left != nil {
					prev = tmp
					tmp = tmp.Left
				}
				root.Val = tmp.Val
				prev.Left = tmp.Right
			}
		}
	}
	
	return root
}
// @lc code=end

