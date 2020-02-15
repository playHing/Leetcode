/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	l := 0
	for i, j := 0, 0; j < len(s); j++ {
		if previ, seen := m[s[j]]; seen && i <= previ {
			i = previ + 1
		}
		m[s[j]] = j
		if l < j-i+1 {
			l = j - i + 1
		}
	}
	return l
}

// @lc code=end

