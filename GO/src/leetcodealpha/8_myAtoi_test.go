package leetcodealpha

import (
	"github.com/stretchr/testify/assert"
	"leetcode"
	"math"
	"testing"
)

func Test8_return_41(t *testing.T) {
	s := "41"
	expect := 41
	result := leetcode.MyAtoi(s)
	assert.Equal(t, expect, result)
}

func Test8_return_first_space_41(t *testing.T) {
	s := "   41"
	expect := 41
	result := leetcode.MyAtoi(s)
	assert.Equal(t, expect, result)
}

func Test8_return_last_space_41(t *testing.T) {
	s := "41   "
	expect := 41
	result := leetcode.MyAtoi(s)
	assert.Equal(t, expect, result)
}

func Test8_return_space_41(t *testing.T) {
	s := "   41   "
	expect := 41
	result := leetcode.MyAtoi(s)
	assert.Equal(t, expect, result)
}

func Test8_return_positive_41(t *testing.T) {
	s := "   +41   "
	expect := 41
	result := leetcode.MyAtoi(s)
	assert.Equal(t, expect, result)
}

func Test8_return_negative_41(t *testing.T) {
	s := "-41   "
	expect := -41
	result := leetcode.MyAtoi(s)
	assert.Equal(t, expect, result)
}

func Test8_return_negative_0(t *testing.T) {
	s := "- 41   "
	expect := 0
	result := leetcode.MyAtoi(s)
	assert.Equal(t, expect, result)
}

func Test8_return_negative_123(t *testing.T) {
	s := "-123-41 123  "
	expect := -123
	result := leetcode.MyAtoi(s)
	assert.Equal(t, expect, result)
}

func Test8_return_positive_123(t *testing.T) {
	s := "123-41 123  "
	expect := 123
	result := leetcode.MyAtoi(s)
	assert.Equal(t, expect, result)
}

func Test8_return_negative_negative_0(t *testing.T) {
	s := "--41 123  "
	expect := 0
	result := leetcode.MyAtoi(s)
	assert.Equal(t, expect, result)
}

func Test8_return_pot_negative_0(t *testing.T) {
	s := ".-41 123  "
	expect := 0
	result := leetcode.MyAtoi(s)
	assert.Equal(t, expect, result)
}

func Test8_return_negative_4_pot(t *testing.T) {
	s := "-4.1 123  "
	expect := -4
	result := leetcode.MyAtoi(s)
	assert.Equal(t, expect, result)
}

func Test8_return_positive_pot_4_0(t *testing.T) {
	s := " .4.1 123  "
	expect := 0
	result := leetcode.MyAtoi(s)
	assert.Equal(t, expect, result)
}

func Test8_return_space_0(t *testing.T) {
	s := "   "
	expect := 0
	result := leetcode.MyAtoi(s)
	assert.Equal(t, expect, result)
}

func Test8_return_negative_negative_negative_0(t *testing.T) {
	s := "----"
	expect := 0
	result := leetcode.MyAtoi(s)
	assert.Equal(t, expect, result)
}

func Test8_return_min_MaxInt32(t *testing.T) {
	s := "9147483648"
	expect := math.MaxInt32
	result := leetcode.MyAtoi(s)
	assert.Equal(t, expect, result)
}

func Test8_return_max_MinInt32(t *testing.T) {
	s := "-9147483648"
	expect := math.MinInt32
	result := leetcode.MyAtoi(s)
	assert.Equal(t, expect, result)
}

func Test8_return_max_over64_MaxInt32(t *testing.T) {
	s := "20000000000000000000"
	expect := math.MaxInt32
	result := leetcode.MyAtoi(s)
	assert.Equal(t, expect, result)
}

func Test8_return_positive_max_over64_MaxInt32(t *testing.T) {
	s := "+20000000000000000000"
	expect := math.MaxInt32
	result := leetcode.MyAtoi(s)
	assert.Equal(t, expect, result)
}

func Test8_return_min_over64_MinInt32(t *testing.T) {
	s := "-20000000000000000000"
	expect := math.MinInt32
	result := leetcode.MyAtoi(s)
	assert.Equal(t, expect, result)
}

func Test8_return_positive_one_0(t *testing.T) {
	s := "+"
	expect := 0
	result := leetcode.MyAtoi(s)
	assert.Equal(t, expect, result)
}

func Test8_return_positive_two_0(t *testing.T) {
	s := "++"
	expect := 0
	result := leetcode.MyAtoi(s)
	assert.Equal(t, expect, result)
}

func Test8_return_negative_one_0(t *testing.T) {
	s := "-"
	expect := 0
	result := leetcode.MyAtoi(s)
	assert.Equal(t, expect, result)
}

func Test8_return_negative_two_0(t *testing.T) {
	s := "--"
	expect := 0
	result := leetcode.MyAtoi(s)
	assert.Equal(t, expect, result)
}

func Test8_return_123456780(t *testing.T) {
	s := "  00000000000123456780"
	expect := 123456780
	result := leetcode.MyAtoi(s)
	assert.Equal(t, expect, result)
}

func Test8_return_positive_123456780(t *testing.T) {
	s := "  +00000000000123456780"
	expect := 123456780
	result := leetcode.MyAtoi(s)
	assert.Equal(t, expect, result)
}

func Test8_return_negative_123456780(t *testing.T) {
	s := "  -00000000000123456780"
	expect := -123456780
	result := leetcode.MyAtoi(s)
	assert.Equal(t, expect, result)
}

func Test8_return_negative_0000000000012345333336780(t *testing.T) {
	s := "  -0000000000021474836410"
	expect := math.MinInt32
	result := leetcode.MyAtoi(s)
	assert.Equal(t, expect, result)
}

func Test8_b11228552307_return_0(t *testing.T) {
	s := " b11228552307"
	expect := 0
	result := leetcode.MyAtoi(s)
	assert.Equal(t, expect, result)
}
