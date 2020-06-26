/*
 * @lc app=leetcode id=127 lang=golang
 *
 * [127] Word Ladder
 */

// @lc code=start
func ladderLength(beginWord string, endWord string, wordList []string) int {
	dict := initDict(wordList)
	queue := []int{-1}
	step, curcnt, nextcnt := 1, 1, 0
	for len(queue) != 0 {
		idx := queue[0]
		queue = queue[1:]

		word := beginWord
		if idx != -1 {
			word = wordList[idx]
		}
		if word == endWord {
			return step
		}

		indexes := idxToAppend(word, dict)
		queue = append(queue, indexes...)
		nextcnt += len(indexes)

		if curcnt == 1 {
			step, curcnt, nextcnt = step+1, nextcnt, 0
		} else {
			curcnt -= 1
		}
	}
	return 0
}

func idxToAppend(word string, dict map[string]int) []int {
	res := make([]int, 0)
	bytes := []byte(word)
	for i := 0; i < len(bytes); i++ {
		ob := bytes[i]
		for c := 0; c < 26; c++ {
			if nb := byte('a'+c); nb != ob {
				bytes[i] = nb
				str := string(bytes)
				if idx, b := dict[str]; b {
					res = append(res, idx)
					delete(dict, str)
				}
			}
		}
		bytes[i] = ob
	}
	return res
}

func initDict(wordList []string) map[string]int {
	dict := make(map[string]int)
	for i, s := range wordList {
		dict[s] = i
	}
	return dict
}
// @lc code=end

