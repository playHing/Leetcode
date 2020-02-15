package goLeetCode

import "strings"

/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 */

// @lc code=start

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	const alpha = "abcdefghijklmnopqrstuvwxyz0123456789"
	i, j := 0, len(s)-1
	for i < j {
		lo := strings.ToLower(string(s[i]))
		if !strings.Contains(alpha, lo) {
			i++
			continue
		}
		hi := strings.ToLower(string(s[j]))
		if !strings.Contains(alpha, hi) {
			j--
			continue
		}
		if lo != hi {
			return false
		} else {
			i++
			j--
		}
	}
	return true
}

// @lc code=end
