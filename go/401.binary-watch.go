/*
 * @lc app=leetcode id=401 lang=golang
 *
 * [401] Binary Watch
 */

// @lc code=start
func readBinaryWatch(num int) []string {
	led := []int{1,2,4,8,1,2,4,8,16,32}
	res := []string{}
	var backtrack func(int, int, []int)
	backtrack = func(idx, n int, time []int) {
		if n == 0 {
			t1, t2 := strconv.Itoa(time[0]), strconv.Itoa(time[1])
			if len(t2) == 1 {
				t2 = "0" + t2
			}
			res = append(res, t1+":"+t2)
			return
		}
		for i := idx; i < len(led); i++ {
			if i < 4 {
				if led[i] + time[0] >= 12 {
					continue
				}
				time[0] += led[i]
				backtrack(i+1, n-1, time)
				time[0] -= led[i]
			} else {
				if led[i] + time[1] >= 60 {
					break
				}
				time[1] += led[i]
				backtrack(i+1, n-1, time)
				time[1] -= led[i]
			}
		}
	}

	backtrack(0, num, []int{0,0})
	return res
}
// @lc code=end

