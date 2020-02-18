/*
 * @lc app=leetcode id=451 lang=golang
 *
 * [451] Sort Characters By Frequency
 */

// @lc code=start
func frequencySort(s string) string {
	freq := make(map[rune]int)
	for _, v := range s {
		if f, b := freq[v]; b {
			freq[v] = f + 1
		} else {
			freq[v] = 1
		}
	}
	maxCnt := 0
	for _, cnt := range freq {
		if cnt > maxCnt {
			maxCnt = cnt
		}
	}
	freqBuckets := make([][]rune, maxCnt)
	for k, cnt := range freq {
		freqBuckets[cnt-1] = append(freqBuckets[cnt-1], k)
	}
	res := make([]rune, len(s))
	pres := 0
	for i := maxCnt - 1; i >= 0; i-- {
		for _, v := range freqBuckets[i] {
			for j := 0; j <= i; j++ {
				res[pres] = v
				pres++
			}
		}
	}
	return string(res)
}

// @lc code=end

