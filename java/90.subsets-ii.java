/*
 * @lc app=leetcode id=90 lang=java
 *
 * [90] Subsets II
 */

// @lc code=start
class Solution {

    List<List<Integer>> res = new LinkedList<>();

    public List<List<Integer>> subsetsWithDup(int[] nums) {
        Arrays.sort(nums);
        backtrack(0, nums, new ArrayList<>());
        return res;
    }

    private void backtrack(int s, int[] nums, List<Integer> tmpList) {
        res.add(new ArrayList<>(tmpList));
        for (int i = s; i < nums.length; ++i) {
            if (i > s && nums[i] == nums[i-1]) continue;
            tmpList.add(nums[i]);
            backtrack(i+1, nums, tmpList);
            tmpList.remove(tmpList.size()-1);
        }
    }
}
// @lc code=end

