/*
 * @lc app=leetcode id=70 lang=java
 *
 * [70] Climbing Stairs
 */

// @lc code=start
class Solution {

    public int climbStairs(int n) {
        if (n <= 0) return 0;
        if (n == 1) return 1;
        int[] stepTable = new int[n];
        stepTable[0] = 1;
        stepTable[1] = 2;
        for (int i = 2; i < n; ++i) {
            stepTable[i] = stepTable[i-1] + stepTable[i-2];
        }
        return stepTable[n-1];
    }
}
// @lc code=end

