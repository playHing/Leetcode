/*
 * @lc app=leetcode id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	k = len(nums) - k
	lo, hi := 0, len(nums)-1
	for lo < hi {
		sep := split(nums, lo, hi)
		if sep < k {
			lo = sep + 1
		} else if sep > k {
			hi = sep - 1
		} else {
			break
		}
	}
	return nums[k]
}

func split(nums []int, lo int, hi int) int {
	i, j := lo, hi+1
	for {
		for i < hi {
			i++
			if !(nums[i] < nums[lo]) {
				break
			}
		}
		for lo < j {
			j--
			if !(nums[lo] < nums[j]) {
				break
			}
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[lo], nums[j] = nums[j], nums[lo]
	return j
}

// @lc code=end

