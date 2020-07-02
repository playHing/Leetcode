/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start
func permute(nums []int) [][]int {
	res := [][]int{}
	if len(nums) == 0 {
		return res
	}
	if len(nums) == 1 {
		return [][]int{[]int{nums[0]}}
	}
	for i := 0; i < len(nums); i++ {
		subarry := make([]int, len(nums)-1)
		copy(subarry[:i], nums[:i])
		copy(subarry[i:], nums[i+1:])
		subpermute := permute(subarry)
		for j := 0; j < len(subpermute); j++ {
			tmp := make([]int, len(subpermute[j])+1)
			tmp[0] = nums[i]
			copy(tmp[1:], subpermute[j])
			subpermute[j] = tmp
		}
		if len(subpermute) > 0 {
			res = append(res, subpermute...)
		}
	}
	return res
}
// @lc code=end

