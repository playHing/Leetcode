/*
 * @lc app=leetcode id=150 lang=golang
 *
 * [150] Evaluate Reverse Polish Notation
 */

// @lc code=start
func evalRPN(tokens []string) int {
    if tokens == nil || len(tokens) == 0 {
		return 0
	}
	sk := make([]int, len(tokens)/2+1)
	skp := -1
	for _, tk := range tokens {
		switch tk {
		case "+", "-", "*", "/":
			val := op(tk, sk[skp-1], sk[skp])
			skp = skp-1
			sk[skp] = val
		default:
			skp++
			sk[skp], _ = strconv.Atoi(tk)
		}
	}
	return sk[0]
}

func op(tk string, v1, v2 int) (val int) {
	switch tk {
	case "+":
		val = v1 + v2
	case "-":
		val = v1 - v2
	case "*":
		val = v1 * v2
	case "/":
		val = v1 / v2
	}
	return
}
// @lc code=end

