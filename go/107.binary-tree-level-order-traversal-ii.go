/*
 * @lc app=leetcode id=107 lang=golang
 *
 * [107] Binary Tree Level Order Traversal II
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
func levelOrderBottom(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	curNum, nextNums := 1, 0
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		tmp := make([]int, curNum)
		for i := 0; i < curNum; i++ {
			cur := queue[i]
			tmp[i] = cur.Val
			if cur.Left != nil {
				queue = append(queue, cur.Left)
				nextNums++
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
				nextNums++
			}
		}
		queue = queue[curNum:]
		res = append(res, tmp)
		curNum, nextNums = nextNums, 0
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
// @lc code=end

