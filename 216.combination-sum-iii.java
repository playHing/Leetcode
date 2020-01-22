/*
 * @lc app=leetcode id=216 lang=java
 *
 * [216] Combination Sum III
 */

// @lc code=start
class Solution {

    List<List<Integer>> res = new LinkedList<>();

    public List<List<Integer>> combinationSum3(int k, int n) {
        if (k <= 0 || n <= 0) return res;
        backtrack(1, k, n, new ArrayList<>());
        return res;
    }

    private void backtrack(int s, int k, int n, List<Integer> tmpList) {
        if (k == 0) {
            if (n == 0) {
                res.add(new ArrayList<>(tmpList));
            }
            return;
        }
        int b = n < 9 ? n : 9;
        for (int i = s; i <= b; ++i) {
            tmpList.add(i);
            backtrack(i+1, k-1, n-i, tmpList);
            tmpList.remove(tmpList.size()-1);
        }
    }
}
// @lc code=end

