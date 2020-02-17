/*
 * @lc app=leetcode id=349 lang=golang
 *
 * [349] Intersection of Two Arrays
 */

// @lc code=start

type void struct{}

var member void

func intersection(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	set := make(map[int]void)
	for _, v := range nums2 {
		set[v] = member
	}
	for _, v := range nums1 {
		_, b := set[v]
		if b {
			delete(set, v)
			res = append(res, v)
		}
	}
	return res
}

// @lc code=end

