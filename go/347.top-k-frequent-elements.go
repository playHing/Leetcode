/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	res := make([]int, 0)
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	cnt := make(map[int][]int)
	cntList := make([]int, len(freq))
	i := 0
	for k, v := range freq {
		cnt[v] = append(cnt[v], k)
		cntList[i] = v
		i++
	}
	i = i-1
	sort.Ints(cntList)
	for k > 0 {
		items := cnt[cntList[i]]
		res = append(res, items...)
		i, k = i-len(items), k-len(items)
	}
	return res
}
// @lc code=end

