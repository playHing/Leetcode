package main

import "sort"

/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	gp := make(map[string][]string)
	for _, s := range strs {
		k := key(s)
		if ss, b := gp[k]; b {
			gp[k] = append(ss, s)
		} else {
			gp[k] = []string{s}
		}
	}
	res := make([][]string, len(gp))
	pres := 0
	for _, g := range gp {
		res[pres] = g
		pres++
	}
	return res
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func key(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

// @lc code=end
