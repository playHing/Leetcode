/*
 * @lc app=leetcode id=77 lang=java
 *
 * [77] Combinations
 */

// @lc code=start
class Solution {

    List<List<Integer>> res = new LinkedList<>();

    public List<List<Integer>> combine(int n, int k) {
        if (n <= 0 || n < k) return res;
        backtrack(1, n, k, new LinkedList<>());
        return res;
    }

    private void backtrack(int s, int n, int k, List<Integer> tmpList) {
        if (k == 0) {
            res.add(new ArrayList<>(tmpList));
            return;
        }
        for (int i = s; k <= n-i+1; ++i) {
            tmpList.add(i);
            backtrack(i+1, n, k-1, tmpList);
            tmpList.remove(tmpList.size()-1);
        }
    }
}
// @lc code=end

