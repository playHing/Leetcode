package main

/*
 * @lc app=leetcode id=149 lang=golang
 *
 * [149] Max Points on a Line
 */

// @lc code=start

type slope struct {
	xl, yl int // relatively prime
}

type line struct {
	slope
}

var noSlope = slope{0, 0}

func maxPoints(points [][]int) int {
	res := 0
	if len(points) == 0 {
		return res
	}

	for i := 0; i < len(points); i++ {
		slopes := make(map[slope]int)
		for j := i + 1; j < len(points); j++ {
			slopes[newSlope(points[i], points[j])]++
		}
		coincide := slopes[noSlope]
		if res < coincide {
			res = coincide
		}
		slopes[noSlope] = 0
		for _, c := range slopes {
			if res < c+coincide {
				res = c + coincide
			}
		}
	}
	return res + 1
}

func newSlope(pi, pj []int) slope {
	yl := pj[1] - pi[1]
	xl := pj[0] - pi[0]
	d := gcd(yl, xl)
	if d != 0 {
		xl, yl = xl/d, yl/d
		if xl < 0 {
			xl, yl = -xl, -yl
		}
		return slope{xl, yl}
	}
	return slope{0, 0}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// @lc code=end
