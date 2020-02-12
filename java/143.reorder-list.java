/*
 * @lc app=leetcode id=143 lang=java
 *
 * [143] Reorder List
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
    public void reorderList(ListNode head) {
        if (head == null || head.next == null) return;
        ListNode prev = null, slow = head, fast = head;
        while (fast != null && fast.next != null) {
            prev = slow;
            slow = slow.next;
            fast = fast.next.next;
        }
        prev.next = null;
        ListNode revHead = reverseList(slow);
        ListNode be4Head = new ListNode(0), ptr = be4Head;

        boolean flag = true;
        while (head != null && revHead != null) {
            if (flag) {
                ptr.next = head;
                head = head.next;
            } else {
                ptr.next = revHead;
                revHead = revHead.next;
            }
            ptr = ptr.next;
            flag = !flag;
        }
        if (head == null) ptr.next = revHead;
        if (revHead == null) ptr.next = head;
    }

    private ListNode reverseList(ListNode head) {
        if (head == null || head.next == null) return head;
        ListNode prev = null, curr = head, next = head.next;
        while (curr != null) {
            curr.next = prev;
            prev = curr;
            curr = next;
            if (next != null) next = next.next;
        }
        return prev;
    }

    private void tester(ListNode head) {
        int itr = 0;
        while (itr++ < 10 && head != null) {
            System.out.print(head.val + "->");
            head = head.next;
        }
        System.out.println();
    }
}

