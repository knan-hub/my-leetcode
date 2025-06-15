package main

import "strings"

// 给定一个非空的字符串 s ，检查是否可以通过由它的一个子串重复多次构成。

func repeatedSubstringPattern(s string) bool {
	// 方法一
	// ss := (s + s)[1 : 2*len(s)-1]
	// return strings.Contains(ss, s)

	// 方法二
	n := len(s)
	for i := 1; i <= n/2; i++ {
		if n%i != 0 {
			continue
		}
		pattern := s[:i]
		if strings.Repeat(pattern, n/i) == s {
			return true
		}
	}
	return false
}
