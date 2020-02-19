/*
 * @lc app=leetcode id=447 lang=golang
 *
 * [447] Number of Boomerangs
 */

// @lc code=start
func numberOfBoomerangs(points [][]int) int {
	res := 0
	for i := 0; i < len(points); i++ {
		cnt := make(map[int]int)
		for j := 0; j < len(points); j++ {
			if j == i {
				continue
			}
			cnt[d(points[i], points[j])]++
		}
		for _, c := range cnt {
			res += c * (c - 1)
		}
	}
	return res
}

func d(p, q []int) int {
	i, j := p[0]-q[0], p[1]-q[1]
	return i*i + j*j
}

// @lc code=end

