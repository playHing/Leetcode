/*
 * @lc app=leetcode id=167 lang=java
 *
 * [167] Two Sum II - Input array is sorted
 */
class Solution {
    public int[] twoSum(int[] nums, int target) {
        if (nums == null) return null;
        int i = 0, j = nums.length-1;
        while (true) {
            if (nums[i] + nums[j] == target) {
                return new int[] {i+1, j+1};
            } else if (nums[i] + nums[j] < target) {
                i++;
            } else {
                j--;
            }
        }
    }
}

