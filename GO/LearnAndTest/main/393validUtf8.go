package main

import "fmt"

func test393(data []int, expect bool) {
	result := validUtf8(data)
	if result == expect {
		fmt.Println(data, "==", result, "succ")
	} else {
		fmt.Println(data, "==", result, "!=", expect, "failed")
	}

}

func main() {
	test393([]int{197, 130}, true)
	test393([]int{197, 130, 1}, true)
	test393([]int{197, 130, 1, 197, 130, 1, 130}, false)
	test393([]int{235, 140, 4}, false)
	test393([]int{235, 140, 4}, false)
	test393([]int{1}, true)                                                                // 1个
	test393([]int{127}, true)                                                              // 1个
	test393([]int{195, 129}, true)                                                         // 2个
	test393([]int{228, 141, 131}, true)                                                    // 3个
	test393([]int{242, 191, 190, 176}, true)                                               // 4个
	test393([]int{242, 191, 190, 176, 65, 12, 99, 111, 228, 141, 131, 195, 129, 58}, true) // 4个
	test393([]int{242, 191, 190, 11}, false)
	test393([]int{242, 191, 190}, false)
	test393([]int{242, 191, 190, 176, 242, 191, 190, 1, 2, 3}, false)
}

const subBytePrefix = 0xc0

func getNumOfBit(i int) int {
	if (i & 0x80) == 0 {
		return 1
	}
	if (i & 0xe0) == 0xc0 {
		return 2
	}
	if (i & 0xf0) == 0xe0 {
		return 3
	}
	if (i & 0xf8) == 0xf0 {
		return 4
	}
	return 0
}

func validUtf8(data []int) bool {
	dLen := len(data)
	for i := 0; i < dLen; {
		num := getNumOfBit(data[i])
		if num <= 0 || i+num > dLen {
			return false
		}
		for j := 1; j < num; j = j + 1 {
			if (data[i+j] & subBytePrefix) != 0x80 {
				return false
			}
		}
		i += num
	}

	return true
}
