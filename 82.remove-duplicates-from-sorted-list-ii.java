/*
 * @lc app=leetcode id=82 lang=java
 *
 * [82] Remove Duplicates from Sorted List II
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
        ListNode be4Head = new ListNode(0);
        be4Head.next = head;
        ListNode prev = be4Head, ptr = head, nextPtr = ptr.next;
        boolean initSaveVal = false;
        int savedVal = 0;
        while (nextPtr != null) {
            if (ptr.val == nextPtr.val || (initSaveVal && ptr.val == savedVal)) {
                initSaveVal = true;
                savedVal = ptr.val;
                prev.next = nextPtr;
            } else {
                prev = ptr;
            }
            ptr = nextPtr;
            nextPtr = nextPtr.next;
        }
        if (initSaveVal && ptr.val == savedVal) prev.next = null;
        return be4Head.next;
    }
}

