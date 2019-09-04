/*
 * @lc app=leetcode id=220 lang=java
 *
 * [220] Contains Duplicate III
 */
class Solution {
    public boolean containsNearbyAlmostDuplicate(int[] nums, int k, int t) {
        if (nums == null) return false;
        TreeSet<Long> ts = new TreeSet<>();
        for (int i = 0; i < nums.length; ++i) {
            if (i >= k+1) ts.remove(new Long(nums[i-k-1]));
            Long ceilingOfLowerBound = ts.ceiling(new Long(nums[i]) - new Long(t));
            if (ceilingOfLowerBound != null && 
                ceilingOfLowerBound <= new Long(nums[i]) + new Long(t)) {
                return true;
            }
            ts.add(new Long(nums[i]));
        }
        return false;
    }
}

