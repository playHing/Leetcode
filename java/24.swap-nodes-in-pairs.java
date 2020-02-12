/*
 * @lc app=leetcode id=24 lang=java
 *
 * [24] Swap Nodes in Pairs
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
    public ListNode swapPairs(ListNode head) {
        if (head == null || head.next == null) return head;
        ListNode be4Head = new ListNode(0);
        be4Head.next = head;
        ListNode prevPtr = be4Head, curPtr = prevPtr.next, nextPtr = curPtr.next;
        while (nextPtr != null) {
            prevPtr.next = nextPtr;
            curPtr.next = nextPtr.next;
            nextPtr.next = curPtr;
            prevPtr = curPtr;
            curPtr = curPtr.next;
            if (curPtr == null) break;
            nextPtr = curPtr.next;
        }
        return be4Head.next;
    }
}

