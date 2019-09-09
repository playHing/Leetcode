/*
 * @lc app=leetcode id=445 lang=java
 *
 * [445] Add Two Numbers II
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

    ListNode res = null;

    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        if (l1 == null) return l2;
        if (l2 == null) return l1;
        int len1 = getLength(l1);
        int len2 = getLength(l2);
        if (len2 > len1) {
            ListNode tmpLN = l1;
            l1 = l2;
            l2 = tmpLN;
            int tmpLen = len1;
            len1 = len2;
            len2 = tmpLen;
        }
        ListNode l1ptr = l1;
        for (int i = len1; i > len2; i--) l1ptr = l1ptr.next;
        int initExt = extOracle(l1ptr, l2);
        int finalExt = extOracle(l1, l1ptr, initExt);
        if (finalExt == 1) {
            ListNode tmp = new ListNode(1);
            tmp.next = res;
            res = tmp;
        }
        return res;
    }

    private int extOracle(ListNode l1, ListNode l1ptr, int initExt) {
        if (l1 == l1ptr) return initExt;
        int ext = extOracle(l1.next, l1ptr, initExt);
        int sum = l1.val + ext;
        ListNode tmp = new ListNode(sum % 10);
        tmp.next = res;
        res = tmp;
        return sum / 10;
    }

    private int extOracle(ListNode l1, ListNode l2) {
        if (l1.next == null && l2.next == null) {
            int sum = l1.val + l2.val;
            res = new ListNode(sum % 10);
            return sum / 10;
        }
        int sum = l1.val + l2.val + extOracle(l1.next, l2.next);
        ListNode tmp = new ListNode(sum % 10);
        tmp.next = res;
        res = tmp;
        return sum / 10;
    }

    private int getLength(ListNode l) {
        int len = 0;
        while (l != null) {
            l = l.next;
            len++;
        }
        return len;
    }
}

