/*
 * @lc app=leetcode id=437 lang=golang
 *
 * [437] Path Sum III
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
func pathSum(root *TreeNode, sum int) int {
	return pathSumRec(root, 0, sum, false)
}

func pathSumRec(root *TreeNode, acc, sum int, started bool) int {
	cnt := 0
    if root == nil {
		return cnt
	}
	acc += root.Val
	if acc == sum {
		cnt += 1
	}
	cnt += pathSumRec(root.Left, acc, sum, true) + pathSumRec(root.Right, acc, sum, true)
	if !started {
		cnt += pathSumRec(root.Left, 0, sum, false) + pathSumRec(root.Right, 0, sum, false)
	}
	return cnt
}
// @lc code=end

