import java.util.List;

/*
 * @lc app=leetcode id=40 lang=java
 *
 * [40] Combination Sum II
 */

// @lc code=start
class Solution {

    List<List<Integer>> res = new LinkedList<>();

    public List<List<Integer>> combinationSum2(int[] candidates, int target) {
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
        int prev = -1;
        for (int i = start; i < candidates.length && candidates[i] <= target; ++i) {
            if (candidates[i] == prev) {
                continue;
            }
            prev = candidates[i];
            tmpList.add(candidates[i]);
            backtrack(i+1, candidates, target-candidates[i], tmpList);
            tmpList.remove(tmpList.size()-1);

        }
    }

}
// @lc code=end

