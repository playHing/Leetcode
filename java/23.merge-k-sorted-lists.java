/*
 * @lc app=leetcode id=23 lang=java
 *
 * [23] Merge k Sorted Lists
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
    public ListNode mergeKLists(ListNode[] lists) {
        if (lists == null || lists.length == 0) return null;
        PriorityQueue<ListNode> queue = new PriorityQueue<>((n1,n2) -> n1.val - n2.val);
        for (ListNode node : lists) {
            if (node != null) queue.offer(node);
        }
        ListNode res = new ListNode(0), ptr = res;
        while (!queue.isEmpty()) {
            ListNode node = queue.poll();
            if (node.next != null) queue.offer(node.next);
            ptr.next = node;
            ptr = node;
        }
        return res.next;
    }
}

