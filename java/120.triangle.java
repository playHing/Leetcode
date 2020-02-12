/*
 * @lc app=leetcode id=120 lang=java
 *
 * [120] Triangle
 */

// @lc code=start
class Solution {

    public int minimumTotal(List<List<Integer>> triangle) {
        if (triangle == null || triangle.size() == 0) return 0;
        int[] memo = new int[triangle.get(triangle.size()-1).size()];
        memo[0] = triangle.get(0).get(0);
        for (int i = 1; i < triangle.size(); ++i) {
            List<Integer> curRow = triangle.get(i);
            int prev = memo[0];
            memo[0] = prev + curRow.get(0);
            for (int j = 1; j < curRow.size()-1; ++j) {
                int tmp = Math.min(prev, memo[j]) + curRow.get(j);
                prev = memo[j];
                memo[j] = tmp;
            }
            memo[curRow.size()-1] = prev + curRow.get(curRow.size()-1);
        }
        int min = memo[0];
        for (int i = 1; i < memo.length; ++i) {
            if (memo[i] < min) min = memo[i];
        }
        return min;
    }
}
// @lc code=end

