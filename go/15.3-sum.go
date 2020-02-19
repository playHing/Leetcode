package main

import "sort"

/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start

type void struct{}

var member void

func threeSum(nums []int) [][]int {
	sort.Sort(sort.IntSlice(nums))
	res := make([][]int, 0)
	for i := 0; i < len(nums); {
		a := nums[i]
		j, k := i+1, len(nums)-1
		for j < k {
			b, c := nums[j], nums[k]
			sum := a + b + c
			if sum == 0 {
				res = append(res, []int{a, b, c})
			}
			if sum >= 0 {
				for j < k && nums[k] == c {
					k--
				}
			}
			if sum <= 0 {
				for j < k && nums[j] == b {
					j++
				}
			}
		}
		for i < len(nums) && nums[i] == a {
			i++
		}
	}
	return res
}

// @lc code=end
