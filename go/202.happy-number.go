/*
 * @lc app=leetcode id=202 lang=golang
 *
 * [202] Happy Number
 */

// @lc code=start

type void struct{}

var member void

func isHappy(n int) bool {
	set := make(map[int]void)
	for {
		sum := 0
		for n != 0 {
			r := n % 10
			sum += r * r
			n /= 10
		}
		if sum == 1 {
			return true
		}
		if _, b := set[sum]; b {
			return false
		} else {
			set[sum] = member
		}
		n = sum
	}
}

// @lc code=end

