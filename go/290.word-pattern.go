package main

import (
	"strings"
)

/*
 * @lc app=leetcode id=290 lang=golang
 *
 * [290] Word Pattern
 */

// @lc code=start

type void struct{}

var member void

func wordPattern(pattern string, str string) bool {
	var lookup [26]string
	strs := strings.Split(str, " ")
	if len(pattern) != len(strs) {
		return false
	}
	for i, s := range strs {
		li := int(pattern[i]) - int('a')
		if p := lookup[li]; p != "" {
			if p != s {
				return false
			}
		} else {
			lookup[li] = s
		}
	}
	set := make(map[string]void)
	for _, s := range lookup {
		if s != "" {
			if _, b := set[s]; b {
				return false
			} else {
				set[s] = member
			}
		}
	}
	return true
}

// @lc code=end
