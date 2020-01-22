/*
 * @lc app=leetcode id=78 lang=java
 *
 * [78] Subsets
 */

// @lc code=start
class Solution {

    List<List<Integer>> res = new LinkedList<>();

    public List<List<Integer>> subsets(int[] nums) {
        if (nums.length == 0) return res;
        backtrack(0, nums, new ArrayList<>());
        return res;
    }

    private void backtrack(int s, int[] nums, List<Integer> tmpList) {
        res.add(new ArrayList<>(tmpList));
        for (int i = s; i < nums.length; ++i) {
            tmpList.add(nums[i]);
            backtrack(i+1, nums, tmpList);
            tmpList.remove(tmpList.size()-1);
        }
    }
}
// @lc code=end

