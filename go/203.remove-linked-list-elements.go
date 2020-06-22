/*
 * @lc app=leetcode id=203 lang=golang
 *
 * [203] Remove Linked List Elements
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	ahead := &ListNode{0, head}
	for prev, cur := ahead, head ;cur != nil; cur = cur.Next {
		if cur.Val == val {
			prev.Next = cur.Next
		} else {
			prev.Next = cur
			prev = cur
		}
	}
	return ahead.Next
}
// @lc code=end

