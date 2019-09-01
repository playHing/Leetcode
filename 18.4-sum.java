/*
 * @lc app=leetcode id=18 lang=java
 *
 * [18] 4Sum
 */
class Solution {
    public List<List<Integer>> fourSum(int[] nums, int target) {
        if (nums == null || nums.length < 4) return new LinkedList<>();
        Arrays.sort(nums);
        List<List<Integer>> res = new LinkedList<>();
        for (int t2 = nums.length - 1; t2 >= 0; --t2) {
            if (t2 < nums.length - 1 && nums[t2] == nums[t2+1]) continue;
            int vt2 = nums[t2];
            for (int t1 = t2 - 1; t1 >= 0; --t1) {
                if (t1 < t2 - 1 && nums[t1] == nums[t1+1]) continue;
                int vt1 = nums[t1];
                int i = 0, j = t1 - 1;
                while (i < j) {
                    int vi = nums[i], vj = nums[j];
                    int sum = vi + vj + vt1 + vt2;
                    if (sum == target) res.add(Arrays.asList(vi, vj, vt1, vt2));
                    if (sum <= target) while (++i < j && nums[i] == vi);
                    if (sum >= target) while (i < --j && nums[j] == vj);
                }
            }
        }
        return res;
    }
}

