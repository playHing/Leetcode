/*
 * @lc app=leetcode id=217 lang=java
 *
 * [217] Contains Duplicate
 */
class Solution {
    public boolean containsDuplicate(int[] nums) {
        if (nums == null) return false;
        Set<Integer> hist = new HashSet<>();
        for (int item : nums) {
            if (hist.contains(item)) return true;
            hist.add(item);
        }
        return false;
    }
}

