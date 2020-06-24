/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
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
	prev, cur := ahead, head
	for cur != nil {
		if cur.Next == nil || cur.Next.Val != cur.Val {
			prev, cur = cur, cur.Next
		} else {
			val := cur.Val
			for cur != nil && cur.Val == val {
				cur = cur.Next
			}
			prev.Next = cur
		}
	}
	return ahead.Next
}
// @lc code=end

