/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    if head == nil || head.Next == nil {
		return true
	}
	ahead := &ListNode{0, head}
	slow, fast := head, head.Next
	prev := ahead
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		
		tmp := slow.Next
		slow.Next = prev
		prev = slow
		slow = tmp
	}

	rightHead := slow.Next
	if fast.Next != nil {
		rightHead = rightHead.Next
	}
	slow.Next = prev
	leftHead := slow

	ahead.Next.Next = nil

	for leftHead != nil {
		if leftHead.Val != rightHead.Val {
			return false
		}
		leftHead, rightHead = leftHead.Next, rightHead.Next
	}
	return true
}
// @lc code=end

