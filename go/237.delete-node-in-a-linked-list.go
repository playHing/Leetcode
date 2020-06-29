/*
 * @lc app=leetcode id=237 lang=golang
 *
 * [237] Delete Node in a Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
    for {
		node.Val = node.Next.Val
		if node.Next.Next == nil {
			node.Next = nil
			break
		}
		node = node.Next
	}
	return
}
// @lc code=end

