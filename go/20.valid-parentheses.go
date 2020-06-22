/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
	match := map[byte]byte{')': '(', ']': '[', '}': '{'}
	stack := make([]byte, 0)
    for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '[', '{':
			stack = append(stack, s[i])
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != match[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
			return false
		}
	}
	return len(stack) == 0
}
// @lc code=end

