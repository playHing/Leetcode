/*
 * @lc app=leetcode id=376 lang=golang
 *
 * [376] Wiggle Subsequence
 */

// @lc code=start
func wiggleMaxLength(nums []int) int {
	max := 1
	n := len(nums)
	if n == 0 {
		return 0
	}
	updp := make([]int, n)
	downdp := make([]int, n)
	
	for i := 0; i < n; i++ {
		updp[i], downdp[i] = 1, 1
		for j := 0; j < i; j++ {
			if nums[i] == nums[j] {
				continue
			}
			if nums[i] < nums[j] {
				if downdp[i] < updp[j] + 1 {
					downdp[i] = updp[j] + 1
					if max < downdp[i] {
						max = downdp[i]
					}
				}
			} else {
				if updp[i] < downdp[j] + 1 {
					updp[i] = downdp[j] + 1
					if max < updp[i] {
						max = updp[i]
					}
				}
			}
		}
	}

	return max
}
// @lc code=end

