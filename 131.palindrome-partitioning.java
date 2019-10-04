/*
 * @lc app=leetcode id=131 lang=java
 *
 * [131] Palindrome Partitioning
 */

// @lc code=start
class Solution {

    boolean[][] dp;

    List<List<String>> res = new LinkedList<>();

    public List<List<String>> partition(String s) {
        if (s != null) {
            initPalindromeIndicator(s);
            partitioning(new Stack<>(), 0, s);
        }
        return res;
    }

    private void partitioning(Stack<String> stack, int start, String s) {
        if (start == s.length()) {
            res.add(new ArrayList<>(stack));
            return;
        }
        for (int i = start; i < s.length(); ++i) {
            if (isPalindrome(start, i)) {
                stack.push(s.substring(start, i+1));
                partitioning(stack, i+1, s);
                stack.pop();
            }
        }
    }

    private void initPalindromeIndicator(String s) {
        if (s.length() == 0) return;
        dp = new boolean[s.length()][s.length()];
        for (int i = 0; i < s.length(); ++i) {
            for (int j = 0; j <= i; ++j) {
                if (s.charAt(i) == s.charAt(j) && (i - j <= 2 || dp[j+1][i-1])) dp[j][i] = true;
            }
        }
    }

    private boolean isPalindrome(int begin, int end) {
        return dp[begin][end];
    }
}
// @lc code=end

