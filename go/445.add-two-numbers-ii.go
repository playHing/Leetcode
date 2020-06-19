/*
 * @lc app=leetcode id=445 lang=golang
 *
 * [445] Add Two Numbers II
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
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	s1, s2 := make([]int, 0), make([]int, 0)
	for ;l1 != nil; l1 = l1.Next {
		s1 = append(s1, l1.Val)
	}
	for ;l2 != nil; l2 = l2.Next {
		s2 = append(s2, l2.Val)
	}
	head := new(ListNode)
	carry := 0
	for i := 1; i <= len(s1) || i <= len(s2); i++ {
		v := carry
		if idx := len(s1)-i; idx >= 0 {
			v += s1[idx]
		}
		if idx := len(s2)-i; idx >= 0 {
			v += s2[idx]
		}
		tmp := &ListNode{v, head.Next}
		head.Next = tmp
		if tmp.Val >= 10 {
			tmp.Val -= 10
			carry = 1
		} else {
			carry = 0
		}
	}
	if carry == 1 {
		tmp := &ListNode{1, head.Next}
		head.Next = tmp
	}
	return head.Next
}

// @lc code=end

