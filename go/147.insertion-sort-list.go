/*
 * @lc app=leetcode id=147 lang=golang
 *
 * [147] Insertion Sort List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	ahead := &ListNode{0, head}
	head, head.Next = head.Next, nil
OUTER:
	for head != nil {
		prev, cur := ahead, ahead.Next
		for cur != nil {
			if head.Val < cur.Val {
				prev.Next = head
				head, head.Next = head.Next, cur
				continue OUTER
			}
			prev, cur = cur, cur.Next
		}
		prev.Next = head
		head, head.Next = head.Next, nil
	}
	return ahead.Next
}
// @lc code=end

