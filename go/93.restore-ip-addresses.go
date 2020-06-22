/*
 * @lc app=leetcode id=93 lang=golang
 *
 * [93] Restore IP Addresses
 */

// @lc code=start
func restoreIpAddresses(s string) []string {
    return tryseperate(s, 4)
}

func tryseperate(s string, ndots int) []string {
	res := []string{}
	for i := 1; i <= len(s) && i <= 3; i++ {
		if i > 1 && s[:1] == "0" {
			break;
		}
		if val, _ := strconv.Atoi(s[:i]); val <= 255 {
			if ndots == 1 {
				if i == len(s) {
					res = append(res, s)
				}
			} else {
				offsets := tryseperate(s[i:], ndots-1)
				for _, offset := range offsets {
					res = append(res, s[:i]+"."+offset)
				}
			}
		}
	}
	return res
}
// @lc code=end

