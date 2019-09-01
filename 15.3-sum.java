/*
 * @lc app=leetcode id=15 lang=java
 *
 * [15] 3Sum
 */
class Solution {
    public List<List<Integer>> threeSum(int[] nums) {
        if (nums == null || nums.length < 3) return new LinkedList<>();
        List<List<Integer>> res = new LinkedList<>();
        Arrays.sort(nums);
        for (int t = nums.length - 1; t >= 0; --t) {
            if (t < nums.length - 1 && nums[t] == nums[t+1]) continue;
            int i = 0, j = t - 1, vt = nums[t];
            while (i < j) {
                int vi = nums[i], vj = nums[j];
                int sum = vi + vj + vt;
                if (sum == 0) res.add(Arrays.asList(vi, vj, vt));
                if (sum <= 0) while (i < j && nums[++i] == vi);
                if (sum >= 0) while (i < j && nums[--j] == vj);
            }
        }
        return res;
    }
}

