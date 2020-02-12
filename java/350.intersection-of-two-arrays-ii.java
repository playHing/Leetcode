/*
 * @lc app=leetcode id=350 lang=java
 *
 * [350] Intersection of Two Arrays II
 */
class Solution {

    // Assume a large amount of repeat elements
    public int[] intersect(int[] nums1, int[] nums2) {
        if (nums1 == null || nums2 == null) return new int[] {};
        Map<Integer, Integer> m1 = new HashMap<>();
        List<Integer> l = new LinkedList<>();
        for (int i = 0; i < nums1.length; ++i) {
            m1.put(nums1[i], 1 + m1.getOrDefault(nums1[i], 0));
        }
        for (int i = 0; i < nums2.length; ++i) {
            int count = m1.getOrDefault(nums2[i], 0);
            if (count == 0) continue;
            l.add(nums2[i]);
            if (count == 1) m1.remove(nums2[i]);
            else m1.put(nums2[i], count-1);
        }
        int[] res = new int[l.size()];
        for (int i = 0; i < res.length; ++i) {
            res[i] = l.remove(0);
        }
        return res;
    }
}

