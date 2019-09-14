/*
 * @lc app=leetcode id=234 lang=java
 *
 * [234] Palindrome Linked List
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
    public boolean isPalindrome(ListNode head) {
        if (head == null || head.next == null) return true;
        ListNode prev = null, slow = head, next = head.next, fast = head;
        while (fast != null && fast.next != null) {
            fast = fast.next.next;
            slow.next = prev;
            prev = slow;
            slow = next;
            next = next.next;
        }
        if (fast != null) slow = next; // odd list
        while (slow != null) {
            if (slow.val != prev.val) return false;
            slow = slow.next;
            prev = prev.next;
        }
        return true;
    }
}

