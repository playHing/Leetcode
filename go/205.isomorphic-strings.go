/*
 * @lc app=leetcode id=205 lang=golang
 *
 * [205] Isomorphic Strings
 */

// @lc code=start

func isIsomorphic(s string, t string) bool {
	return isInjective(&s, &t) && isInjective(&t, &s)
}

func isInjective(ps *string, pt *string) bool {
	s, t := *ps, *pt
	mapping := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if v, b := mapping[t[i]]; b {
			if v != s[i] {
				return false
			}
		} else {
			mapping[t[i]] = s[i]
		}
	}
	return true
}

// @lc code=end

