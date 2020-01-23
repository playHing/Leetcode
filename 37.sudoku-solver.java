/*
 * @lc app=leetcode id=37 lang=java
 *
 * [37] Sudoku Solver
 */

// @lc code=start
class Solution {

    char[][] board = null;

    public void solveSudoku(char[][] board) {
        this.board = board;
        dfs(board, 0);
    }

    private boolean dfs(char[][] board, int loc) {
        if (loc == 81) return true;
        int i = loc/9, j = loc%9;
        if (board[i][j] != '.') {
            return dfs(board, loc+1);
        }
        boolean[] flag = new boolean[9];
        if (!isValid(i, j, flag)) return false;
        for (int k = 0; k < 9; ++k) {
            if (flag[k]) {
                board[i][j] = (char)('1' + k);
                if (dfs(board, loc+1)) return true;
            }
        }
        board[i][j] = '.';
        return false;
    }

    private boolean isValid(int i, int j, boolean[] flag) {
        Arrays.fill(flag, true);
        for (int k = 0; k < 9; ++k) {
            if (board[i][k] != '.') flag[board[i][k]-'1'] = false;
            if (board[k][j] != '.') flag[board[k][j]-'1'] = false;
            int r = i/3 * 3 + k/3, c = j/3 * 3 + k%3;
            if (board[r][c] != '.') flag[board[r][c]-'1'] = false;
        }
        for (int k = 0; k < 9; ++k) {
            if (flag[k]) return true;
        }
        return false;
    }
}
// @lc code=end

