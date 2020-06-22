/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */

// @lc code=start
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	board := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	res := []string{""}
	for i := 0; i < len(digits); i++ {
		d := int(digits[i])-int('0')
		l := board[d]
		curlen := len(res)
		for _, s := range res {
			for j := 0; j < len(l); j++ {
					res = append(res, s + string(l[j]))
			}
		}
		res = res[curlen:]
	}
	return res
}
// @lc code=end

