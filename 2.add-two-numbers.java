/*
 * @lc app=leetcode id=2 lang=java
 *
 * [2] Add Two Numbers
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
    
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        if (l1 == null) return l2;
        if (l2 == null) return l1;
        int ext = 0;
        ListNode head = new ListNode(-1), ptr = head;
        while (l1 != null || l2 != null) {
            int sum = ext + getVal(l1) + getVal(l2);
            ptr.next = new ListNode(sum % 10);
            ext = sum / 10;
            l1 = next(l1);
            l2 = next(l2);
            ptr = ptr.next;
        }
        if (ext == 1) ptr.next = new ListNode(1);
        return head.next;
    }

    private int getVal(ListNode l) {
        return l != null ? l.val : 0;
    }

    private ListNode next(ListNode l) {
        return l != null ? l.next : null;
    }

}

