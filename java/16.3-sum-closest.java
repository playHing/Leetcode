/*
 * @lc app=leetcode id=16 lang=java
 *
 * [16] 3Sum Closest
 */
class Solution {
    public int threeSumClosest(int[] nums, int target) {
        if (nums == null || nums.length < 3) return 0;
        Arrays.sort(nums);
        int closet = nums[0] + nums[1] + nums[2];
        for (int k = nums.length - 1; k >= 0; --k) {
            if (k < nums.length - 1 && nums[k] == nums[k+1]) continue;
            int vk = nums[k];
            int i = 0, j = k - 1;
            while (i < j) {
                int vi = nums[i], vj = nums[j];
                int sum = vi + vj + vk;
                if (d(sum, target) < d(closet, target)) closet = sum;
                if (sum <= target) while (++i < j && nums[i] == vi);
                if (sum >= target) while (i < --j && nums[j] == vj);
            }
        }
        return closet;
    }

    private int d(int a, int b) {
        return a < b ? b - a : a - b;
    }
}

