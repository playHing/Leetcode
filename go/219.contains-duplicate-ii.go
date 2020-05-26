/*
 * @lc app=leetcode id=219 lang=golang
 *
 * [219] Contains Duplicate II
 */

// @lc code=start

func containsNearbyDuplicate(nums []int, k int) bool {
	cache := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if _, b := cache[nums[i]]; b {
			return true
		}
		cache[nums[i]] = struct{}{}
		if i >= k {
			delete(cache, nums[i-k])
		}
	}
	return false
}

// @lc code=end

