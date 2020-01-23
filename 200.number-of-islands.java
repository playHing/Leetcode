/*
 * @lc app=leetcode id=200 lang=java
 *
 * [200] Number of Islands
 */

// @lc code=start
class Solution {
    public int numIslands(char[][] grid) {
        if (grid == null || grid.length == 0 || grid[0].length == 0) return 0;
        int res = 0;
        int[][] moves = new int[][] {{-1,0},{0,1},{1,0},{0,-1}};
        boolean[][] visited = new boolean[grid.length][grid[0].length];
        for (int i = 0; i < grid.length; ++i) {
            for (int j = 0; j < grid[0].length; ++j) {
                if (visited[i][j] || grid[i][j] != '1') continue;
                res++;
                dfs(visited, i, j, moves, grid);
            }
        }
        return res;
    }

    private void dfs(boolean[][] visited, int i, int j, int[][] moves, char[][] grid) {
        visited[i][j] = true;
        for (int[] move : moves) {
            int ni = i+move[0];
            int nj = j+move[1];
            if (0 <= ni && ni < grid.length && 
                0 <= nj && nj < grid[0].length && 
                grid[ni][nj] == '1' && !visited[ni][nj]) {
                dfs(visited, ni, nj, moves, grid);
            }
        }
    }
}
// @lc code=end

