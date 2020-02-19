package main

import "sort"

/*
 * @lc app=leetcode id=18 lang=golang
 *
 * [18] 4Sum
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Sort(sort.IntSlice(nums))
	for i := 0; i < len(nums); {
		a := nums[i]
		for j := i + 1; j < len(nums); {
			b := nums[j]
			k, l := j+1, len(nums)-1
			for k < l {
				c, d := nums[k], nums[l]
				sum := a + b + c + d
				if sum == target {
					res = append(res, []int{a, b, c, d})
				}
				if sum >= target {
					for k < l && nums[l] == d {
						l--
					}
				}
				if sum <= target {
					for k < l && nums[k] == c {
						k++
					}
				}
			}
			for j < len(nums) && nums[j] == b {
				j++
			}
		}
		for i < len(nums) && nums[i] == a {
			i++
		}
	}
	return res
}

// @lc code=end
