/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */

// @lc code=start
func longestPalindrome(s string) string {
	if len(s) == 0 {
        return ""
    }
    dp := make([][]bool, len(s))
    for i := range dp {
        dp[i] = make([]bool, len(s))
        dp[i][i] = true
    }
    max, mi, mj := 1, 0, 0
    for i := 0; i+1 < len(s); i++ {
        dp[i][i+1] = s[i] == s[i+1]
        if dp[i][i+1] && max < 2 {
            max, mi, mj = 1, i, i+1
        }
    }
    for l := 2; l < len(s); l++ {
        for i := 0; i+l < len(s); i++ {
            dp[i][i+l] = s[i] == s[i+l] && dp[i+1][i+l-1]
            if dp[i][i+l] && max < l+1 {
                max, mi, mj = l+1, i, i+l
            }
        }
    }
    return s[mi:mj+1]
}
// @lc code=end

