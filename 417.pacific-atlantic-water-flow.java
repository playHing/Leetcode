/*
 * @lc app=leetcode id=417 lang=java
 *
 * [417] Pacific Atlantic Water Flow
 */

// @lc code=start
class Solution {

    static final int[][] moves = new int[][] {{-1,0},{0,1},{1,0},{0,-1}};
    List<List<Integer>> res = new LinkedList<>();

    public List<List<Integer>> pacificAtlantic(int[][] matrix) {
        if (matrix == null || matrix.length == 0 || matrix[0].length == 0) return res;

        boolean[][] pacificReachable = new boolean[matrix.length][matrix[0].length];
        boolean[][] atlanticReachable = new boolean[matrix.length][matrix[0].length];

        for (int i = 0; i < matrix.length; ++i) {
            dfs(matrix, pacificReachable, i, 0);
            dfs(matrix, atlanticReachable, i, matrix[0].length-1);
        }
        for (int j = 0; j < matrix[0].length; ++j) {
            dfs(matrix, pacificReachable, 0, j);
            dfs(matrix, atlanticReachable, matrix.length-1, j);
        }
        for (int i = 0; i < matrix.length; ++i) {
            for (int j = 0; j < matrix[0].length; ++j) {
                if (pacificReachable[i][j] && atlanticReachable[i][j]) {
                    res.add(Arrays.asList(new Integer[]{i, j}));
                }
            }
        }
        return res;
    }

    private void dfs(int[][] matrix, boolean[][] reachable, int i, int j) {
        dfs(matrix, reachable, i, j, matrix[i][j]);
    }

    private void dfs(int[][] matrix, boolean[][] reachable, int i, int j, int toVal) {
        if (matrix[i][j] < toVal || reachable[i][j]) return;
        reachable[i][j] = true;
        for (int[] move : moves) {
            int ni = i + move[0], nj = j + move[1];
            if (ni < 0 || ni >= matrix.length || nj < 0 || nj >= matrix[0].length) {
                continue;
            }
            dfs(matrix, reachable, ni, nj, matrix[i][j]);
        }
    }

}
// @lc code=end

