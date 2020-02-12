/*
 * @lc app=leetcode id=27 lang=java
 *
 * [27] Remove Element
 */
class Solution {
    public int removeElement(int[] nums, int val) {
        if (nums == null || nums.length == 0) return 0;
        int mainIdx = 0;
        for (int runningIdx = 0; runningIdx < nums.length; ++runningIdx) {
            if (nums[runningIdx] != val) {
                if (runningIdx != mainIdx) {
                    nums[mainIdx] = nums[runningIdx];
                }
                mainIdx++;
            }
        }
        return mainIdx;
    }
}

