/*
 * @lc app=leetcode id=86 lang=java
 *
 * [86] Partition List
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
    public ListNode partition(ListNode head, int x) {
        if (head == null || head.next == null) return head;
        ListNode hl = new ListNode(-1), pl = hl;
        ListNode hg = new ListNode(-1), pg = hg;
        while (head != null) {
            if (head.val < x) {
                pl.next = head;
                pl = head;
            } else {
                pg.next = head;
                pg = head;
            }
            head = head.next;
        }
        pl.next = hg.next;
        pg.next = null;
        return hl.next;
    }
}

