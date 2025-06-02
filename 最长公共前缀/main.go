package main

// 编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串 ""。

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	str := strs[0]
	s := make([]byte, 0)
	for i := 0; i < len(str); i++ {
		isMatch := true
		for j := 1; j < len(strs); j++ {
			curStr := strs[j]
			if len(curStr) <= i || curStr[i] != str[i] {
				isMatch = false
				break
			}
		}
		if !isMatch {
			break
		}
		s = append(s, str[i])
	}
	return string(s)
}
