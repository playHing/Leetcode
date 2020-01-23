/*
 * @lc app=leetcode id=79 lang=java
 *
 * [79] Word Search
 */

// @lc code=start
class Solution {

    public boolean exist(char[][] board, String word) {
        if (word == null || word.length() == 0) return true;
        if (board.length == 0 || board[0].length == 0) return false;
        for (int i = 0; i < board.length; ++i) {
            for (int j = 0; j < board[0].length; ++j) {
                if (backstrack(i, j, 0, board, word)) return true;
            }
        }
        return false;
    }

    private boolean backstrack(int i, int j, int c, char[][] board, String word) {
        if (i < 0 || i >= board.length || j < 0 || j >= board[0].length) {
            return false;
        }
        if (board[i][j] == word.charAt(c)) {
            if (c == word.length()-1) return true;
            char tmp = board[i][j];
            board[i][j] = 'x';
            boolean res = backstrack(i+1, j, c+1, board, word) ||
                backstrack(i, j+1, c+1, board, word) ||
                backstrack(i-1, j, c+1, board, word) ||
                backstrack(i, j-1, c+1, board, word);
            board[i][j] = tmp;
            return res;
        } else {
            return false;
        }
    }


}
// @lc code=end

