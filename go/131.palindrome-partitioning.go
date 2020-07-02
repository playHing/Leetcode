/*
 * @lc app=leetcode id=131 lang=golang
 *
 * [131] Palindrome Partitioning
 */

// @lc code=start
func partition(s string) [][]string {
	res := [][]string{}
	
	var backtrack func(int, []string)
	backtrack = func(idx int, pt []string) {
		if idx == len(s) {
			tmp := make([]string, len(pt))
			copy(tmp, pt)
			res = append(res, tmp)
			return
		}
		for i := idx; i < len(s); i++ {
			if isPalindrome(s[idx:i+1]) {
				pt = append(pt, s[idx:i+1])
				backtrack(i+1, pt)
				pt = pt[:len(pt)-1]
			}
		}
	}

	backtrack(0, []string{})
	return res
}

func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
// @lc code=end

