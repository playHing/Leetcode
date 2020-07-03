/*
 * @lc app=leetcode id=236 lang=golang
 *
 * [236] Lowest Common Ancestor of a Binary Tree
 */

// @lc code=start
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pp, pq := []*TreeNode{}, []*TreeNode{}
	dfs(root, p, &pp)
	dfs(root, q, &pq)
	if len(pp) > len(pq) {
		pp, pq = pq, pp
	}
	for i := 1; i < len(pp); i++ {
		if pp[i] != pq[i] {
			return pp[i-1]
		}
	}
	return pp[len(pp)-1]
}

func dfs(root, node *TreeNode, path *[]*TreeNode) bool {
	if root == nil {
		return false
	}
	*path = append(*path, root)
	if root == node {
		return true
	}
	exist := dfs(root.Left, node, path) || dfs(root.Right, node, path)
	if exist {
		return true
	}
	*path = (*path)[:len(*path)-1]
	return false
}
// @lc code=end

