/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	hh := new(ListNode)
	h := hh
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			h.Next = l1
			l1 = l1.Next
		} else {
			h.Next = l2
			l2 = l2.Next
		}
		h = h.Next
	}
	if l1 == nil {
		h.Next = l2
	} else {
		h.Next = l1
	}
	return hh.Next
}

// @lc code=end

