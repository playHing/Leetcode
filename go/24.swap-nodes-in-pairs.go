/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ahead := &ListNode{0, head}
	prev, cur, next := ahead, head, head.Next
	for next != nil {
		tmpNext := next.Next
		prev.Next, cur.Next, next.Next = next, tmpNext, cur
		if tmpNext == nil {
			break
		}
		prev, cur, next = cur, tmpNext, tmpNext.Next
	}
	return ahead.Next
}
// @lc code=end

