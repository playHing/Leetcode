/*
 * @lc app=leetcode id=206 lang=java
 *
 * [206] Reverse Linked List
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
    public ListNode reverseList(ListNode head) {
        if (head == null || head.next == null) return head;
        ListNode nextNode = head.next;
        head.next = null;
        while (nextNode != null) {
            ListNode next2Node = nextNode.next;
            nextNode.next = head;
            head = nextNode;
            nextNode = next2Node;
        }
        return head;
    }
}

