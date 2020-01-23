/*
 * @lc app=leetcode id=130 lang=java
 *
 * [130] Surrounded Regions
 */

// @lc code=start
class Solution {

    int[][] moves = new int[][] {{-1,0},{0,1},{1,0},{0,-1}};

    public void solve(char[][] board) {
        if (board == null || board.length == 0 || board[0].length == 0) return;
        boolean[][] visited = new boolean[board.length][board[0].length];
        for (int i = 0; i < board.length; ++i) {
            dfs(board, visited, i, 0);
            dfs(board, visited, i, board[0].length-1);
        }
        for (int j = 1; j < board[0].length-1; ++j) {
            dfs(board, visited, 0, j);
            dfs(board, visited, board.length-1, j);
        }
        for (int i = 0; i < board.length; ++i) {
            for (int j = 0; j < board[0].length; ++j) {
                if (board[i][j] == 'O' && !visited[i][j]) {
                    board[i][j] = 'X';
                }
            }
        }
    }

    private void dfs(char[][] board, boolean[][] visited, int i, int j) {
        if (board[i][j] != 'O' || visited[i][j]) return;
        visited[i][j] = true;
        for (int[] move : moves) {
            int ni = i + move[0], nj = j + move[1];
            if (ni < 0 || ni >= board.length || nj < 0 || nj >= board[0].length) {
                continue;
            }
            dfs(board, visited, ni, nj);
        }
    }
}
// @lc code=end

