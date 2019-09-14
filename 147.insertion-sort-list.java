/*
 * @lc app=leetcode id=147 lang=java
 *
 * [147] Insertion Sort List
 */
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public ListNode insertionSortList(ListNode head) {
        if (head == null || head.next == null) return head;
        ListNode be4Head = new ListNode(0);
        ListNode prev = be4Head, curr = head, next = null;
        while (curr != null) {
            next = curr.next;
            while (prev.next != null && 
                prev.next.val < curr.val) prev = prev.next;
            curr.next = prev.next;
            prev.next = curr;
            prev = be4Head;
            curr = next;
        }
        return be4Head.next;
    }

}

