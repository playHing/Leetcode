/*
 * @lc app=leetcode id=199 lang=golang
 *
 * [199] Binary Tree Right Side View
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
func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
    if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	curcnt, nextcnt := 1, 0
	for len(queue) != 0 {
		for i := 0; i < curcnt; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
				nextcnt++
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				nextcnt++
			}
		}
		res = append(res, queue[curcnt-1].Val)
		queue = queue[curcnt:]
		curcnt, nextcnt = nextcnt, 0
	}
	return res
}
// @lc code=end

