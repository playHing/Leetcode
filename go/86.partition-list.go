/*
 * @lc app=leetcode id=86 lang=golang
 *
 * [86] Partition List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	l1, l2 := new(ListNode), new(ListNode)
	h1, h2 := l1, l2
	for head != nil {
		tmp := head.Next
		head.Next = nil
		if head.Val < x {
			l1.Next = head
			l1 = l1.Next
		} else {
			l2.Next = head
			l2 = l2.Next
		}
		head = tmp
	}
	l1.Next = h2.Next
	return h1.Next
}
// @lc code=end

