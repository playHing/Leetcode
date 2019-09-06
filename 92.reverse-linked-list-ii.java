/*
 * @lc app=leetcode id=92 lang=java
 *
 * [92] Reverse Linked List II
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
    public ListNode reverseBetween(ListNode head, int m, int n) {
        if (head == null || head.next == null) return head;

        ListNode be4rl = head;
        int i = 1;
        if (m != 1) while (++i < m) be4rl = be4rl.next;
        else {
            be4rl = new ListNode(-1);
            be4rl.next = head;
        }

        ListNode rlp = be4rl.next;
        ListNode rlpNext = rlp.next;
        while (i++ < n) {
            ListNode tmp = rlpNext.next;
            rlpNext.next = rlp;
            rlp = rlpNext;
            rlpNext = tmp;
        }
        be4rl.next.next = rlpNext;
        be4rl.next = rlp;
        return m == 1 ? be4rl.next : head;
    }
}

