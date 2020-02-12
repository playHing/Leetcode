/*
 * @lc app=leetcode id=19 lang=java
 *
 * [19] Remove Nth Node From End of List
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
    public ListNode removeNthFromEnd(ListNode head, int n) {
        ListNode nHead = head, nTail = head;
        while (nTail != null && n-- > 0) nTail = nTail.next;
        if (n != -1) return head.next;
        while (nTail.next != null) {
            nHead = nHead.next;
            nTail = nTail.next;
        }
        if (nHead.next != null) {
            nHead.next = nHead.next.next;
        }
        return head;
    }
}

