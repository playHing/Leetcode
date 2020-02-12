/*
 * @lc app=leetcode id=83 lang=java
 *
 * [83] Remove Duplicates from Sorted List
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
    public ListNode deleteDuplicates(ListNode head) {
        if (head == null || head.next == null) return head;
        ListNode prev = head, ptr = head.next;
        while (ptr != null) {
            if (prev.val != ptr.val) prev = ptr;
            else prev.next = ptr.next;
            ptr = ptr.next;
        }
        return head;
    }
}

