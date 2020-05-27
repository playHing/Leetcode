import "strings"

/*
 * @lc app=leetcode id=71 lang=golang
 *
 * [71] Simplify Path
 */

// @lc code=start
func simplifyPath(path string) string {
	var stack []string
	s := strings.Split(path, "/")
	for _, ss := range s {
		switch ss {
		case "":
		case ".":
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, ss)
		}
	}
	res := ""
	for _, ss := range stack {
		res += "/" + ss
	}
	if res == "" {
		res = "/"
	}
	return res
}

// @lc code=end

