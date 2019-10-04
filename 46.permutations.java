/*
 * @lc app=leetcode id=46 lang=java
 *
 * [46] Permutations
 */

// @lc code=start
class Solution {

    List<List<Integer>> res = new LinkedList<>();

    public List<List<Integer>> permute(int[] nums) {
        if (nums != null && nums.length > 0) {
            backtrack(1, nums, new LinkedList<Integer>() {{add(nums[0]);}} );
        }
        return res;
    }

    private void backtrack(int idx, int[] nums, List<Integer> tmpList) {
        if (idx == nums.length) {
            res.add(new ArrayList<>(tmpList));
            return;
        }
        for (int i = 0; i <= tmpList.size(); ++i) {
            tmpList.add(i, nums[idx]);
            backtrack(idx+1, nums, tmpList);
            tmpList.remove(i);
        }
    }



}
// @lc code=end

