package leetcodealpha

import (
	"github.com/stretchr/testify/assert"
	"leetcode"
	"testing"
)

func Test14_2_no(t *testing.T) {
	s := []string{"LVIII", "123"}
	expect := ""
	result := leetcode.LongestCommonPrefix(s)
	assert.Equal(t, expect, result)
}

func Test14_1_b(t *testing.T) {
	s := []string{"b"}
	expect := "b"
	result := leetcode.LongestCommonPrefix(s)
	assert.Equal(t, expect, result)
}

func Test14_1(t *testing.T) {
	s := []string{"LVIII"}
	expect := "LVIII"
	result := leetcode.LongestCommonPrefix(s)
	assert.Equal(t, expect, result)
}

func Test14_1_(t *testing.T) {
	s := []string{""}
	expect := ""
	result := leetcode.LongestCommonPrefix(s)
	assert.Equal(t, expect, result)
}

func Test14_2_LV(t *testing.T) {
	s := []string{"LVIII", "LV"}
	expect := "LV"
	result := leetcode.LongestCommonPrefix(s)
	assert.Equal(t, expect, result)
}

func Test14_2_LVxx(t *testing.T) {
	s := []string{"LVIII", "LVxx"}
	expect := "LV"
	result := leetcode.LongestCommonPrefix(s)
	assert.Equal(t, expect, result)
}

func Test14_2_2LVxx(t *testing.T) {
	s := []string{"LVxx", "LVxx"}
	expect := "LVxx"
	result := leetcode.LongestCommonPrefix(s)
	assert.Equal(t, expect, result)
}

func Test14_3_LVxx(t *testing.T) {
	s := []string{"LVIII", "LVxx", "LVxx123"}
	expect := "LV"
	result := leetcode.LongestCommonPrefix(s)
	assert.Equal(t, expect, result)
}

func Test14_3_(t *testing.T) {
	s := []string{"LVIII", "LVxx", ""}
	expect := ""
	result := leetcode.LongestCommonPrefix(s)
	assert.Equal(t, expect, result)
}

func Test14_3_no(t *testing.T) {
	s := []string{"LVIII", "xaLVxx", "asfd"}
	expect := ""
	result := leetcode.LongestCommonPrefix(s)
	assert.Equal(t, expect, result)
}

func Test14_3_fl(t *testing.T) {
	s := []string{"flower", "flow", "flight"}
	expect := "fl"
	result := leetcode.LongestCommonPrefix(s)
	assert.Equal(t, expect, result)
}

func Test14_5_fl(t *testing.T) {
	s := []string{"flower", "flow", "flight", "flight", "flight"}
	expect := "fl"
	result := leetcode.LongestCommonPrefix(s)
	assert.Equal(t, expect, result)
}

func Test14_5_fl_(t *testing.T) {
	s := []string{"flight", "flight", "flight", "flower", "flow"}
	expect := "fl"
	result := leetcode.LongestCommonPrefix(s)
	assert.Equal(t, expect, result)
}

func Test14_5_987(t *testing.T) {
	s := []string{"987654321", "98765432", "9876543", "98765", "987"}
	expect := "987"
	result := leetcode.LongestCommonPrefix(s)
	assert.Equal(t, expect, result)
}

func Test14_5_9876(t *testing.T) {
	s := []string{"987654321", "98765432", "98765", "9876", "98765", "98765432"}
	expect := "9876"
	result := leetcode.LongestCommonPrefix(s)
	assert.Equal(t, expect, result)
}

func Test14_5_(t *testing.T) {
	s := []string{"987654321", "", "98765", "9876", "98765", "98765432"}
	expect := ""
	result := leetcode.LongestCommonPrefix(s)
	assert.Equal(t, expect, result)
}

func Test14_5_9(t *testing.T) {
	s := []string{"987654321", "98765", "98ac5", "9er6", "9ert5", "98765ery2"}
	expect := "9"
	result := leetcode.LongestCommonPrefix(s)
	assert.Equal(t, expect, result)
}

func Test14_2_b(t *testing.T) {
	s := []string{"", "b"}
	expect := ""
	result := leetcode.LongestCommonPrefix(s)
	assert.Equal(t, expect, result)
}
