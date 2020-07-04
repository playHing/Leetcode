/*
 * @lc app=leetcode id=61 lang=golang
 *
 * [61] Rotate List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil {
		return head
	}
	len := 1
	for p := head; p.Next != nil; {
		len, p = len+1, p.Next
	}
	if len <= k {
		k = k % len
	}
	if k == 0 {
		return head
	}
	ahead := &ListNode{0, head}
	slow, fast := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow, fast = slow.Next, fast.Next
	}
	fast.Next = ahead.Next
	ahead.Next = slow.Next
	slow.Next = nil
	return ahead.Next
}
// @lc code=end

