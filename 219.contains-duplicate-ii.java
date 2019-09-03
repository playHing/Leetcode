/*
 * @lc app=leetcode id=219 lang=java
 *
 * [219] Contains Duplicate II
 */
class Solution {
    public boolean containsNearbyDuplicate(int[] nums, int k) {
        if (nums == null) return false;
        Set<Integer> lookup = new HashSet<>();
        for (int r = 0; r < nums.length; ++r) {
            if (r-k-1 >= 0) lookup.remove(nums[r-k-1]);
            if (lookup.contains(nums[r])) return true;
            lookup.add(nums[r]);
        }
        return false;
    }
}

