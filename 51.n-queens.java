/*
 * @lc app=leetcode id=51 lang=java
 *
 * [51] N-Queens
 */

// @lc code=start
class Solution {

    List<List<String>> res = new LinkedList<>();
        boolean[] r1;
        boolean[] r2;
        boolean[] r3;

    public List<List<String>> solveNQueens(int n) {
        if (n <= 0) {
            res.add(new ArrayList<>());
            return res;
        }
        r1 = new boolean[n];
        r2 = new boolean[2*n-1];
        r3 = new boolean[2*n-1];
        char[] ts = new char[n];
        for (int i = 0; i < n; ++i) ts[i] = '.';
        backtrack(0, new LinkedList<>(), n, ts);
        return res;
    }

    private void backtrack(int i, List<String> tmpList, int n, char[] ts) {
        if (i == n) {
            res.add(new ArrayList<>(tmpList));
            return;
        }
        for (int j = 0; j < n; ++j) {
            if (!r1[j] && !r2[i+j] && !r3[(j-i)+n-1]) {
                r1[j] = true;
                r2[i+j] = true;
                r3[(j-i)+n-1] = true;
                ts[j] = 'Q';
                tmpList.add(new String(ts));
                ts[j] = '.';
                backtrack(i+1, tmpList, n, ts);
                tmpList.remove(tmpList.size()-1);
                r1[j] = false;
                r2[i+j] = false;
                r3[(j-i)+n-1] = false;

            }
        }
    }



}
// @lc code=end

