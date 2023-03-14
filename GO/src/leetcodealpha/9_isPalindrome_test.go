package leetcodealpha

import (
	"github.com/stretchr/testify/assert"
	"leetcode"
	"math"
	"testing"
)

func Test9_return_121(t *testing.T) {
	s := 121
	expect := true
	result := leetcode.IsPalindrome(s)
	assert.Equal(t, expect, result)
}

func Test9_return_1213(t *testing.T) {
	s := 1213
	expect := false
	result := leetcode.IsPalindrome(s)
	assert.Equal(t, expect, result)
}

func Test9_return_1(t *testing.T) {
	s := 1
	expect := true
	result := leetcode.IsPalindrome(s)
	assert.Equal(t, expect, result)
}

func Test9_return_12(t *testing.T) {
	s := 12
	expect := false
	result := leetcode.IsPalindrome(s)
	assert.Equal(t, expect, result)
}
func Test9_return_12321(t *testing.T) {
	s := 12321
	expect := true
	result := leetcode.IsPalindrome(s)
	assert.Equal(t, expect, result)
}

func Test9_return_000(t *testing.T) {
	s := 000
	expect := true
	result := leetcode.IsPalindrome(s)
	assert.Equal(t, expect, result)
}

func Test9_return_11(t *testing.T) {
	s := 11
	expect := true
	result := leetcode.IsPalindrome(s)
	assert.Equal(t, expect, result)
}

func Test9_return_max(t *testing.T) {
	s := math.MaxInt
	expect := false
	result := leetcode.IsPalindrome(s)
	assert.Equal(t, expect, result)
}
