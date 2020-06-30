/*
 * @lc app=leetcode id=257 lang=golang
 *
 * [257] Binary Tree Paths
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
func binaryTreePaths(root *TreeNode) []string {
    if root == nil {
		return []string{}
	}
	return bfs(root, "")
}

func bfs(root *TreeNode, path string) []string {
	if path == "" {
		path = strconv.Itoa(root.Val)
	} else {
		path = path + "->" + strconv.Itoa(root.Val)
	}
	if root.Left == nil && root.Right == nil {
		return []string{path}
	}
	res := []string{}
	if root.Left != nil {
		res = append(res, bfs(root.Left, path)...)
	}
	if root.Right != nil {
		res = append(res, bfs(root.Right, path)...)
	}
	return res
}
// @lc code=end

