/*
 * @lc app=leetcode id=203 lang=java
 *
 * [203] Remove Linked List Elements
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
    public ListNode removeElements(ListNode head, int val) {
        if (head == null) return null;
        ListNode be4Head = new ListNode(0), prev = be4Head;
        be4Head.next = head;
        while (head != null) {
            if (head.val == val) prev.next = head.next; 
            else prev = head;
            head = head.next;
        }
        return be4Head.next;
    }
}

