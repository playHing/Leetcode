/*
 * @lc app=leetcode id=1 lang=java
 *
 * [1] Two Sum
 */
class Solution {
    public int[] twoSum(int[] nums, int target) {
        if (nums == null) return new int[] {};
        Map<Integer, Integer> idxMap = new HashMap<>();
        for (int i = 0; i < nums.length; ++i) {
            if (idxMap.containsKey(nums[i]))
                return new int[] {idxMap.get(nums[i]), i};
            else
                idxMap.put(target-nums[i], i);
        }
        return new int[] {};
    }
}

