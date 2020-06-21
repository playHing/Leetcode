/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    if head == nil || head.Next == nil {
		return head
	}
	ahead := &ListNode{0, head}
	prev, cur, next := ahead, head, head.Next
	for i := 1; i < m; i++ {
		prev, cur, next = cur, next, next.Next
	}
	prevRef, curRef := prev, cur
	for i := m; i <= n; i++ {
		cur.Next = prev
		if next != nil {
			prev, cur, next = cur, next, next.Next
		} else {
			prevRef.Next = cur
			curRef.Next = next
			return ahead.Next
		}
	}
	prevRef.Next = prev
	curRef.Next = cur
	return ahead.Next
}
// @lc code=end

