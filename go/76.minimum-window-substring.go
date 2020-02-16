package main

/*
 * @lc app=leetcode id=76 lang=golang
 *
 * [76] Minimum Window Substring
 */

// @lc code=start
func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}
	var cp, mem [256]int
	for i := 0; i < len(t); i++ {
		cp[int(t[i])]++
		mem[int(s[i])]++
	}
	var res string
	for j, i := 0, len(t)-1; ; {
		if isContaining(&mem, &cp) {
			if len(res) == 0 || len(res) > i+1-j {
				res = s[j : i+1]
			}
			mem[int(s[j])]--
			j++
		} else {
			i++
			if i >= len(s) {
				break
			}
			mem[int(s[i])]++
		}
	}
	return res
}

func isContaining(mem *[256]int, cp *[256]int) bool {
	for i := 0; i < 256; i++ {
		if (*mem)[i] < (*cp)[i] {
			return false
		}
	}
	return true
}

// @lc code=end
