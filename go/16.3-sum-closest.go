package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	sort.Sort(sort.IntSlice(nums))
	res := nums[0] + nums[1] + nums[2]
	md := d(res, target)
	for i := 0; i < len(nums); i++ {
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			d := d(sum, target)
			if d < md {
				md, res = d, sum
				if md == 0 {
					return res // = target
				}
			}
			if sum < target {
				j++
			} else {
				k--
			}
		}
	}
	return res
}

func d(sum, target int) int {
	d := sum - target
	if d > 0 {
		return d
	} else {
		return -d
	}
}

// @lc code=end
