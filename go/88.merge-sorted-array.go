/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	r1 := m - 1
	r2 := n - 1
	for r := m + n - 1; r >= 0 && r1 >= 0 && r2 >= 0; r-- {
		if nums1[r1] > nums2[r2] {
			nums1[r] = nums1[r1]
			r1--
		} else {
			nums1[r] = nums2[r2]
			r2--
		}
	}
	for r2 >= 0 {
		nums1[r2] = nums2[r2]
		r2--
	}
}

// @lc code=end

