/*
 * @lc app=leetcode id=88 lang=java
 *
 * [88] Merge Sorted Array
 */
class Solution {
    public void merge(int[] nums1, int m, int[] nums2, int n) {
        if (nums1 == null || nums2 == null) return;
        int i = m-1, j = n-1, k = m+n-1;
        while (k > i) {
            if (i == -1) {
                for (;k >= 0;) nums1[k--] = nums2[j--];
                break;
            }
            nums1[k--] = nums1[i] < nums2[j] ? nums2[j--] : nums1[i--];
        }
    }
}

