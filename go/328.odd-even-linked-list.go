/*
 * @lc app=leetcode id=328 lang=golang
 *
 * [328] Odd Even Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
		return head
	}
	oddahead, evenahead := new(ListNode), new(ListNode)
	cur, oddhead, evenhead := head, oddahead, evenahead
	for i := 1; cur != nil; i, cur = i+1, cur.Next {
		if i % 2 == 1 {
			oddhead.Next = cur
			oddhead = cur
		} else {
			evenhead.Next = cur
			evenhead = cur
		}
	}
	evenhead.Next = nil
	oddhead.Next = evenahead.Next
	return oddahead.Next
}
// @lc code=end

