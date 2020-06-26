/*
 * @lc app=leetcode id=126 lang=golang
 *
 * [126] Word Ladder II
 */

// @lc code=start

// "aaa"\n"bbb"\n["aab","baa","bab","bbb"]

type Item struct {
	Tail int
	Path []int
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	
	res := make([][]string, 0)

	dict := initDict(wordList)
	queue := []Item{{-1, []int{-1}}}

	for len(queue) != 0 && len(res) == 0 {
		curcnt := len(queue)
		indexesToDelete := make(map[int]struct{})

		for j := 0; j < curcnt; j++ {
			item := queue[j]
			idx := item.Tail

			word := beginWord
			if idx != -1 {
				word = wordList[idx]
			}
			if word == endWord {
				path := make([]string, len(item.Path))
				path[0] = beginWord
				for k := 1; k < len(item.Path); k++ {
					path[k] = wordList[item.Path[k]]
				}
				res = append(res, path)
			}

			indexes := idxToAppend(word, dict)
			items := make([]Item, len(indexes))
			for k := 0; k < len(indexes); k++ {
				items[k] = Item{indexes[k], append([]int{}, item.Path...)}
				items[k].Path = append(items[k].Path, indexes[k])
				indexesToDelete[indexes[k]] = struct{}{}
			}
			queue = append(queue, items...)

		}
		queue = queue[curcnt:]
		for k := range indexesToDelete {
			delete(dict, wordList[k])
		}
	}
	return res
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

