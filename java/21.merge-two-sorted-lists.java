/*
 * @lc app=leetcode id=21 lang=java
 *
 * [21] Merge Two Sorted Lists
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
    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        if (l1 == null) return l2;
        if (l2 == null) return l1;
        ListNode be4Head = new ListNode(0), ptr = be4Head;
        while (true) {
            if (l1 == null) {
                ptr.next = l2;
                break;
            }
            if (l2 == null) {
                ptr.next = l1;
                break;
            }
            if (l1.val < l2.val) {
                ptr.next = l1;
                ptr = l1;
                l1 = l1.next;
            } else {
                ptr.next = l2;
                ptr = l2;
                l2 = l2.next;
            }
        }
        return be4Head.next;
    }
}

