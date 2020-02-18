/*
 * @lc app=leetcode id=454 lang=golang
 *
 * [454] 4Sum II
 */

// @lc code=start
func fourSumCount(A []int, B []int, C []int, D []int) int {
	// O(n^2)
	cnt := make(map[int]int)
	for _, va := range A {
		for _, vb := range B {
			if v, b := cnt[va+vb]; b {
				cnt[va+vb] = v + 1
			} else {
				cnt[va+vb] = 1
			}
		}
	}
	res := 0
	for _, vc := range C {
		for _, vd := range D {
			if v, b := cnt[-vc-vd]; b {
				res += v
			}
		}
	}
	return res
}

// @lc code=end

