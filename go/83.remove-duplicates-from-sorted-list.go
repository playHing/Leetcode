/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ahead := &ListNode{0, head}
	prev, cur, next := ahead, head, head.Next
	lastSeen := cur.Val - 1
	for {
		for lastSeen == cur.Val {
			prev.Next = next
			if next == nil {
				break
			}
			cur, next = next, next.Next
		}
		lastSeen = cur.Val
		if next == nil {
			break
		}
		prev, cur, next = cur, next, next.Next
	}
	return ahead.Next
}
// @lc code=end

