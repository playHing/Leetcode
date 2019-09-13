/*
 * @lc app=leetcode id=25 lang=java
 *
 * [25] Reverse Nodes in k-Group
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
    public ListNode reverseKGroup(ListNode head, int k) {
        
        if (head == null || head.next == null) return head;
        ListNode be4Head = new ListNode(0);
        be4Head.next = head;
        boolean init = false;
        ListNode be4Itr = be4Head, firstItemAtItr = be4Itr.next;
        ListNode prev = be4Head, cur = prev.next, nex = cur.next;

        while (firstItemAtItr != null) {
            int itr = k;
            outer:
            for (int dump = 1; dump <= 2; ++dump) {

                while (itr > 0) {
                    if (nex == null && itr > 1) {
                        nex = prev;
                        prev = null;
                        itr = k - itr + 1;
                        continue outer;
                    }
                    cur.next = prev;
                    prev = cur;
                    cur = nex;
                    if (nex != null) nex = nex.next;
                    itr--;
                }

                if (dump == 2) return be4Head.next;
                break;
            }

            be4Itr.next = prev;
            if (!init) {
                be4Head.next = prev;
                init = true;
            }
            firstItemAtItr.next = cur;
            be4Itr = firstItemAtItr;
            firstItemAtItr = cur;
            prev = be4Itr;
        }
        
        return be4Head.next;
    }
}

