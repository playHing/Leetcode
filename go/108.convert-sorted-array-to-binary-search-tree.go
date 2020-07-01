/*
 * @lc app=leetcode id=108 lang=golang
 *
 * [108] Convert Sorted Array to Binary Search Tree
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
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
    if len(nums) == 1 {
		return &TreeNode{nums[0], nil, nil}
	}
	if len(nums) == 2 {
		return &TreeNode{nums[1], &TreeNode{nums[0], nil, nil}, nil}
	}
	return &TreeNode{nums[len(nums)/2], 
		sortedArrayToBST(nums[:len(nums)/2]), 
		sortedArrayToBST(nums[len(nums)/2+1:])}

}
// @lc code=end

