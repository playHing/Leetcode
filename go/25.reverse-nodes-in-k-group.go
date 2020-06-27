/*
 * @lc app=leetcode id=25 lang=golang
 *
 * [25] Reverse Nodes in k-Group
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    if head == nil || k <= 1 {
		return head
	}
	ahead := &ListNode{0, head}
	prev, cur, next := ahead, head, head.Next
	OUTER:
	for {
		tmp := cur
		for i := 0; i < k; i++ {
			if tmp == nil {
				break OUTER
			}
			tmp = tmp.Next
		}
		firstPrev := prev
		for i := 0; i < k; i++ {
			cur.Next = prev
			prev, cur = cur, next
			if next != nil {
				next = next.Next
			}
		}
		tmp = prev
		prev = firstPrev.Next
		firstPrev.Next.Next = cur
		firstPrev.Next = tmp
	}
	return ahead.Next
}
// @lc code=end

