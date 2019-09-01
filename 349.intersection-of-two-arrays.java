/*
 * @lc app=leetcode id=349 lang=java
 *
 * [349] Intersection of Two Arrays
 */
class Solution {

    // Assume a large amount of repeat elements
    public int[] intersection(int[] nums1, int[] nums2) {
        if (nums1 == null || nums2 == null) return new int[] {};
        Set<Integer> s1 = new HashSet<>();
        List<Integer> l = new LinkedList<>();
        for (int i = 0; i < nums1.length; ++i) s1.add(nums1[i]);
        for (int i = 0; i < nums2.length; ++i) {
            if (s1.contains(nums2[i])) {
                l.add(nums2[i]);
                s1.remove(nums2[i]);
            }
        }
        int[] res = new int[l.size()];
        for (int i = 0; i < res.length; ++i) {
            res[i] = l.remove(0);
        }
        return res;
    }
}

