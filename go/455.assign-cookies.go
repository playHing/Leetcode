/*
 * @lc app=leetcode id=455 lang=golang
 *
 * [455] Assign Cookies
 */

// @lc code=start
func findContentChildren(g []int, s []int) int {
	num := 0
	sort.Ints(g)
	sort.Ints(s)

	for childIdx, cookieIdx :=len(g)-1, len(s)-1; childIdx >= 0 && cookieIdx >= 0; {
		if s[cookieIdx] < g[childIdx] {
			childIdx--
		} else {
			cookieIdx--
			childIdx--
			num++
		}
	}
	return num
}
// @lc code=end

