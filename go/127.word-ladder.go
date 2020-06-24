/*
 * @lc app=leetcode id=127 lang=golang
 *
 * [127] Word Ladder
 */

// @lc code=start
func ladderLength(beginWord string, endWord string, wordList []string) int {
	visited := make([]bool, len(wordList))
	queue := make([]int, 0, len(wordList))
	queue = append(queue, -1)
	step, curcnt, nextcnt := 1, 1, 0
	for len(queue) != 0 {
		idx, word := queue[0], beginWord
		queue = queue[1:]
		if idx != -1 {
			visited[idx] = true
			word = wordList[idx]
		}
		if word == endWord {
			return step
		}
		for i := 0; i < len(wordList); i++ {
			if !visited[i] && dist(wordList[i], word) == 1 {
				queue = append(queue, i)
				nextcnt++
			}
		}
		if curcnt == 1 {
			step, curcnt, nextcnt = step+1, nextcnt, 0
		} else {
			curcnt -= 1
		}
	}
	return 0
}

func dist(w1, w2 string) (d int) {
	for i := 0; i < len(w1); i++ {
		if w1[i] != w2[i] {
			d++
			if d >= 2 {
				return 2
			}
		}
	}
	return d
}
// @lc code=end

