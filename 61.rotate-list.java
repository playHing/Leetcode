/*
 * @lc app=leetcode id=61 lang=java
 *
 * [61] Rotate List
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
    public ListNode rotateRight(ListNode head, int k) {
        if (head == null || head.next == null) return head;
        int cnt = k, len = 0;
        ListNode kHead = head, kTail = head;
        while (kTail != null && cnt-- > 0) {
            kTail = kTail.next;
            len++;
        }
        if (cnt != -1) {
            cnt = k % len;
            kTail = head;
            while (kTail != null && cnt-- > 0) kTail = kTail.next;
        }
        while (kTail.next != null) {
            kHead = kHead.next;
            kTail = kTail.next;
        }
        kTail.next = head;
        head = kHead.next;
        kHead.next = null;
        return head;
    }
}

