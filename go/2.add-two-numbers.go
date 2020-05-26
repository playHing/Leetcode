/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	ptr := head
	r := 0

	for l1 != nil || l2 != nil || r > 0 {
		next := new(ListNode)
		next.Val = r

		if l1 != nil {
			next.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			next.Val += l2.Val
			l2 = l2.Next
		}

		r = next.Val / 10
		next.Val %= 10
		ptr.Next = next
		ptr = next
	}

	return head.Next
}

// @lc code=end

