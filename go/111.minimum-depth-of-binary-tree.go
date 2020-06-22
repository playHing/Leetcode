/*
 * @lc app=leetcode id=111 lang=golang
 *
 * [111] Minimum Depth of Binary Tree
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
func minDepth(root *TreeNode) int {
    if root == nil {
		return 0
	}
	d, curcnt, nextcnt := 1, 1, 0
	queue := []*TreeNode{root}
	for {
		for i := 0; i < curcnt; i++ {
			node := queue[i]
			l, r := node.Left, node.Right
			if l == nil && r == nil {
				return d
			}
			if l != nil {
				queue = append(queue, l)
				nextcnt++
			}
			if r != nil {
				queue = append(queue, r)
				nextcnt++
			}
		}
		queue = queue[curcnt:]
		d, curcnt, nextcnt = d+1, nextcnt, 0
	}
}
// @lc code=end

