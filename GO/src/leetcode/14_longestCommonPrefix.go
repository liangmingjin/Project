package leetcode

func LongestCommonPrefix(strs []string) string {
	ret := len(strs[0])
	j := ret
	for i := 0; i < len(strs); i++ {
		if ret > len(strs[i]) {
			ret = len(strs[i])
		}
		if ret <= 0 {
			return ""
		}
		for j = 0; j < ret; j++ {
			if strs[0][j] != strs[i][j] {
				break
			}
		}
		if ret > j {
			ret = j
		}
	}

	return strs[0][0:ret]
}
