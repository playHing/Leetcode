/*
 * @lc app=leetcode id=328 lang=java
 *
 * [328] Odd Even Linked List
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
    public ListNode oddEvenList(ListNode head) {
        if (head == null || head.next == null) return head;
        ListNode oddHead = head, oddPtr = oddHead;
        ListNode evenHead = head.next, evenPtr = evenHead;
        head = evenHead.next;
        boolean isOdd = true;
        while (head != null) {
            if (isOdd) {
                oddPtr.next = head;
                oddPtr = oddPtr.next;
            } else {
                evenPtr.next = head;
                evenPtr = evenPtr.next;
            }
            head = head.next;
            isOdd = !isOdd;
        }
        oddPtr.next = evenHead;
        evenPtr.next = null;
        return oddHead;
    }
}

