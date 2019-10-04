/*
 * @lc app=leetcode id=47 lang=java
 *
 * [47] Permutations II
 */

// @lc code=start
class Solution {
    
    List<List<Integer>> res = new LinkedList<>();
    
    public List<List<Integer>> permuteUnique(int[] nums) {
        Arrays.sort(nums);
        if (nums != null) {
            backtrack(0, -1, nums, new LinkedList<Integer>());
        }
        return res;
    }

    private void backtrack(int idx, int prev, int[] nums, List<Integer> tmpList) {
        if (idx == nums.length) {
            res.add(new ArrayList<>(tmpList));
            return;
        }
        boolean isNextDuplicate = idx+1 != nums.length && nums[idx] == nums[idx+1];
        for (int i = prev+1; i <= tmpList.size(); ++i) {
            tmpList.add(i, nums[idx]);
            backtrack(idx+1, isNextDuplicate ? i : -1, nums, tmpList);
            tmpList.remove(i);
        }
    }

}
// @lc code=end

