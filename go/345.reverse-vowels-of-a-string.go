import "strings"

/*
 * @lc app=leetcode id=345 lang=golang
 *
 * [345] Reverse Vowels of a String
 */

// @lc code=start

const vowels = "aeiouAEIOU"

func reverseVowels(s string) string {
	i, j := 0, len(s)-1
	rs := []rune(s)
OUTER:
	for i < j {
		for !strings.Contains(vowels, string(rs[i])) {
			i++
			if i >= j {
				break OUTER
			}
		}
		for !strings.Contains(vowels, string(rs[j])) {
			j--
			if i >= j {
				break OUTER
			}
		}
		rs[i], rs[j] = rs[j], rs[i]
		i++
		j--
	}
	return string(rs)
}

// @lc code=end

