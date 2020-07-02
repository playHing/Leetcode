/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	backtrack(0, -1, nums, &res, []int{})
	return res
}

func backtrack(idx, prev int, nums []int, res *[][]int, per []int) {
	fmt.Println(per)
	if idx == len(nums) {
		tmp := make([]int, len(per))
		copy(tmp, per)
		*res = append(*res, tmp)
		return
	}
	nextDup := false
	if idx != len(nums)-1 && nums[idx] == nums[idx+1] {
		nextDup = true
	}
	for i := prev+1; i <= len(per); i++ {
		per = append(per, nums[idx])
		copy(per[i+1:], per[i:])
		per[i] = nums[idx]

		nextPrev := -1
		if nextDup {
			nextPrev = i
		}

		backtrack(idx+1, nextPrev, nums, res, per)
		copy(per[i:], per[i+1:])
		per = per[:len(per)-1]
	}
}
// @lc code=end

