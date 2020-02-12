/*
 * @lc app=leetcode id=39 lang=java
 *
 * [39] Combination Sum
 */

// @lc code=start
class Solution {

    List<List<Integer>> res = new LinkedList<>();

    public List<List<Integer>> combinationSum(int[] candidates, int target) {
        if (candidates.length == 0) return res;
        Arrays.sort(candidates);
        backtrack(0, candidates, target, new ArrayList<>());
        return res;
    }

    private void backtrack(int start, int[] candidates, int target, List<Integer> tmpList) {
        if (target == 0) {
            res.add(new ArrayList<>(tmpList));
            return;
        }
        for (int i = start; i < candidates.length && candidates[i] <= target; ++i) {
            tmpList.add(candidates[i]);
            backtrack(i, candidates, target-candidates[i], tmpList);
            tmpList.remove(tmpList.size()-1);
        }
    }
}
// @lc code=end

