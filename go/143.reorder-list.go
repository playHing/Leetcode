/*
 * @lc app=leetcode id=143 lang=golang
 *
 * [143] Reorder List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
	if head == nil || head.Next == nil {
		return
	}
	ptr := head
	list := make([]*ListNode, 0)
    for ptr != nil {
		list = append(list, ptr)
		ptr = ptr.Next
	}
	for i := 0; i < len(list)/2; i++ {
		list[i].Next = list[len(list)-1-i]
		list[len(list)-1-i].Next = list[i+1]
	}
	list[len(list)/2].Next = nil
}
// @lc code=end

