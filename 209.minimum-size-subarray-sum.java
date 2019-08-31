/*
 * @lc app=leetcode id=209 lang=java
 *
 * [209] Minimum Size Subarray Sum
 */
class Solution {
    public int minSubArrayLen(int s, int[] nums) {
        if (nums == null || nums.length == 0) return 0;
        int min = nums.length + 1, sum = 0;
        int i = 0, j = -1;
        while (i < nums.length) {
            if (sum < s && j < nums.length - 1) sum += nums[++j];
            else sum -= nums[i++];
            if (sum >= s && j-i+1 < min) min = j-i+1;
        }
        return min == nums.length + 1 ? 0 : min;
    }
}

