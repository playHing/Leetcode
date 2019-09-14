/*
 * @lc app=leetcode id=237 lang=java
 *
 * [237] Delete Node in a Linked List
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
    public void deleteNode(ListNode node) {
        if (node == null || node.next == null) return;
        ListNode next = node.next;
        while (true) {
            node.val = next.val;
            if (next.next == null) break;
            node = next;
            next = next.next;
        }
        node.next = null;
    }
}

