/*
 * @lc app=leetcode id=148 lang=golang
 *
 * [148] Sort List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	// merge sort
	return MergeSort(head)
}

func MergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	rightPart := MergeSort(slow.Next)
	slow.Next = nil
	leftPart := MergeSort(head)
	return mergeStep(leftPart, rightPart)
}

func mergeStep(lp, rp *ListNode) *ListNode {
	ahead := &ListNode{0, nil}
	cur := ahead
	for {
		if lp != nil && rp != nil {
			if lp.Val < rp.Val {
				cur.Next = lp
				lp = lp.Next
			} else {
				cur.Next = rp
				rp = rp.Next
			}
			cur = cur.Next
		} else {
			if lp != nil {
				cur.Next = lp
			} else if rp != nil {
				cur.Next = rp
			}
			break
		}
	}
	return ahead.Next
}
// @lc code=end

