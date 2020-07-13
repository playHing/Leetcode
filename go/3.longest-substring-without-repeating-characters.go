/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
        return 0
    }
    max, start := 1, 0
    m := make(map[byte]int)
    for i := 0; i < len(s); i++ {
        c := s[i]
        if pos, b := m[c]; b && start <= pos {
            start = pos + 1
        }
        m[c] = i
        
        if max < i - start + 1 {
            max = i - start + 1
        }
    }
    return max
}

// @lc code=end

