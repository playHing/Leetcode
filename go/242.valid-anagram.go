/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	freq := make(map[rune]int)
	rt := []rune(t)
	for _, v := range rt {
		if f, b := freq[v]; b {
			freq[v] = f + 1
		} else {
			freq[v] = 1
		}
	}
	rs := []rune(s)
	for _, v := range rs {
		if f, b := freq[v]; b {
			if f == 0 {
				return false
			}
			freq[v] = f - 1
		} else {
			return false
		}
	}
	return true
}

// @lc code=end

