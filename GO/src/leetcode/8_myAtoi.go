package leetcode

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func change(c byte) string {
	if c == '+' {
		return "\\+[\\d]*"
	}
	return fmt.Sprintf("%c[\\d]*", c)
}

func NumCompare(s1, s2 string) int {
	//优化比较，先比较位数再比较大小
	if len(s1) == len(s2) {
		return strings.Compare(s1, s2)
	} else if len(s1) > len(s2) {
		return +1
	} else {
		return -1
	}
}

func MyAtoi(s string) int {
	TrimPrefixS := strings.Trim(s, " ")
	if len(TrimPrefixS) == 0 {
		fmt.Println("TrimPrefixS == 0")
		return 0
	}
	//解释正则表达式
	reg := regexp.MustCompile(change(TrimPrefixS[0]))
	if reg == nil {
		fmt.Println("MustCompile err")
		return 0
	}
	//提取关键信息
	result := reg.FindStringSubmatch(TrimPrefixS)
	if result == nil {
		fmt.Println("FindStringSubmatch err")
		return 0
	}
	if result[0][0] == '-' {
		if NumCompare(strings.TrimLeft(result[0][1:], "0"), fmt.Sprintf("%d", math.MinInt32)[1:]) > 0 {
			return math.MinInt32
		}
	} else if result[0][0] == '+' {
		if NumCompare(strings.TrimLeft(result[0][1:], "0"), fmt.Sprintf("%d", math.MaxInt32)) > 0 {
			return math.MaxInt32
		}
	} else if result[0][0] >= '0' && result[0][0] <= '9' {
		if NumCompare(strings.TrimLeft(result[0], "0"), fmt.Sprintf("%d", math.MaxInt32)) > 0 {
			return math.MaxInt32
		}
	}
	if i, err := strconv.Atoi(result[0]); err == nil {
		return i
	} else {
		fmt.Println("strconv.Atoi err:", err)
	}
	return 0
}
