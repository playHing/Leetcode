/*
 * @lc app=leetcode id=283 lang=java
 *
 * [283] Move Zeroes
 */
class Solution {

    // Assume nums is sparse
    public void moveZeroes(int[] nums) {
        if (nums == null) return;
        int mainIdx = 0;
        for (int runningIdx = 0; runningIdx < nums.length; runningIdx++) {
            if (nums[runningIdx] != 0) {
                if (runningIdx != mainIdx) {
                    nums[mainIdx] = nums[runningIdx];
                    nums[runningIdx] = 0;
                }
                mainIdx++;
            }
        }
    }
    
}

