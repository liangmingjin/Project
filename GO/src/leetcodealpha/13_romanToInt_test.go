package leetcodealpha

import (
	"github.com/stretchr/testify/assert"
	"leetcode"
	"testing"
)

func Test13_III(t *testing.T) {
	s := "III"
	expect := 3
	result := leetcode.RomanToInt(s)
	assert.Equal(t, expect, result)
}

func Test13_I(t *testing.T) {
	s := "I"
	expect := 1
	result := leetcode.RomanToInt(s)
	assert.Equal(t, expect, result)
}

func Test13_IV(t *testing.T) {
	s := "IV"
	expect := 4
	result := leetcode.RomanToInt(s)
	assert.Equal(t, expect, result)
}

func Test13_VI(t *testing.T) {
	s := "VI"
	expect := 6
	result := leetcode.RomanToInt(s)
	assert.Equal(t, expect, result)
}

func Test13_IX(t *testing.T) {
	s := "IX"
	expect := 9
	result := leetcode.RomanToInt(s)
	assert.Equal(t, expect, result)
}

func Test13_XI(t *testing.T) {
	s := "XI"
	expect := 11
	result := leetcode.RomanToInt(s)
	assert.Equal(t, expect, result)
}

func Test13_XL(t *testing.T) {
	s := "XL"
	expect := 40
	result := leetcode.RomanToInt(s)
	assert.Equal(t, expect, result)
}

func Test13_LXIX(t *testing.T) {
	s := "LXIX"
	expect := 69
	result := leetcode.RomanToInt(s)
	assert.Equal(t, expect, result)
}

func Test13_CDXLIX(t *testing.T) {
	s := "CDXLIX"
	expect := 449
	result := leetcode.RomanToInt(s)
	assert.Equal(t, expect, result)
}

func Test13_MCMXCIV(t *testing.T) {
	s := "MCMXCIV"
	expect := 1994
	result := leetcode.RomanToInt(s)
	assert.Equal(t, expect, result)
}

func Test13_LVIII(t *testing.T) {
	s := "LVIII"
	expect := 58
	result := leetcode.RomanToInt(s)
	assert.Equal(t, expect, result)
}

func Test13_MDCXXXIII(t *testing.T) {
	s := "MDCXXXIII"
	expect := 1633
	result := leetcode.RomanToInt(s)
	assert.Equal(t, expect, result)
}

func Test13_MMMMCMXCIX(t *testing.T) {
	s := "MMMMCMXCIX"
	expect := 4999
	result := leetcode.RomanToInt(s)
	assert.Equal(t, expect, result)
}

func Test13_DXII(t *testing.T) {
	s := "DXII"
	expect := 512
	result := leetcode.RomanToInt(s)
	assert.Equal(t, expect, result)
}
