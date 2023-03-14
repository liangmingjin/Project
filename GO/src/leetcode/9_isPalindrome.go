package leetcode

import "strconv"

func IsPalindrome(x int) bool {
	s := strconv.Itoa(x)
	for i, j := 0, len(s)-1; i <= j+1; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
