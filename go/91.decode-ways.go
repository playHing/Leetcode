/*
 * @lc app=leetcode id=91 lang=golang
 *
 * [91] Decode Ways
 */

// @lc code=start
func numDecodings(s string) int {
    if len(s) == 0 || s[0] == '0' {
        return 0
    }
	dp := make([]int, len(s)+1)
	dp[0], dp[1] = 1, 1
    if s[len(s)-1] == '0' {
        dp[1] = 0
    }
	for i := 2; i <= len(s); i++ {
        if s[len(s)-i] == '0' {
            dp[i] = 0
            if s[len(s)-i-1] > '2' {
                return 0
            }
        } else if s[len(s)-i] == '1' {
			dp[i] = dp[i-1] + dp[i-2]
		} else if s[len(s)-i] == '2' {
			if s[len(s)-i+1] <= '6' {
				dp[i] = dp[i-1] + dp[i-2]
			} else {
				dp[i] = dp[i-1]
			}
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(s)]
}
// @lc code=end

