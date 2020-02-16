package main

/*
 * @lc app=leetcode id=438 lang=golang
 *
 * [438] Find All Anagrams in a String
 */

// @lc code=start
func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	if len(s) < len(p) {
		return res
	}
	const a = int('a')
	var t [26]int
	for i := 0; i < len(p); i++ {
		t[int(p[i])-a] += 1
		t[int(s[i])-a] -= 1
	}
	if isAna(&t) {
		res = append(res, 0)
	}
	for i := len(p); i < len(s); i++ {
		t[int(s[i])-a] -= 1
		t[int(s[i-len(p)])-a] += 1
		if isAna(&t) {
			res = append(res, i-len(p)+1)
		}
	}
	return res
}

func isAna(t *[26]int) bool {
	for i := 0; i < 26; i++ {
		if (*t)[i] != 0 {
			return false
		}
	}
	return true
}

// @lc code=end
