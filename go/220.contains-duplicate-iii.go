/*
 * @lc app=leetcode id=220 lang=golang
 *
 * [220] Contains Duplicate III
 */

// @lc code=start

// TODO: should use Tree set
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	cache := make(map[int]struct{})
	for i := 0; i <= k && i < len(nums); i++ {
		if appendToCache(cache, nums[i], t) {
			return true
		}
	}
	for i := k+1; i < len(nums); i++ {
		delete(cache, nums[i-k-1])
		if appendToCache(cache, nums[i], t) {
			return true
		}
	}
	return false
}

// return true if exist
func appendToCache(cache map[int]struct{}, v, t int) bool {
	for v2 := range cache {
		if isAcceptedDiff(v, v2, t) {
			return true
		}
	}
	cache[v] = struct{}{}
	return false
}

func isAcceptedDiff(v1, v2, t int) bool {
	if v1 < v2 {
		v1, v2 = v2, v1
	}
	return v1 - v2 <= t
}


// @lc code=end

