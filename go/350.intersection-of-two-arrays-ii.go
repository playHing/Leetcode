/*
 * @lc app=leetcode id=350 lang=golang
 *
 * [350] Intersection of Two Arrays II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	freq := make(map[int]int)
	for _, v := range nums2 {
		if f, b := freq[v]; b {
			freq[v] = f + 1
		} else {
			freq[v] = 1
		}
	}
	for _, v := range nums1 {
		if f, b := freq[v]; b {
			res = append(res, v)
			if f == 1 {
				delete(freq, v)
			} else {
				freq[v] = f - 1
			}
		}
	}
	return res
}

// @lc code=end

