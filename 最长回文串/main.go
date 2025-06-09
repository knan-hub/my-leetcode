package main

// 给定一个包含大写字母和小写字母的字符串 s ，返回通过这些字母构造成的最长的回文串的长度。
// 在构造过程中，请注意区分大小写 。比如 "Aa" 不能当做一个回文字符串。

func longestPalindrome(s string) int {
	memo := map[byte]int{}
	arr := []byte(s)
	for _, v := range arr {
		memo[v]++
	}
	length := 0
	hasOdd := false
	for _, v := range memo {
		if v%2 == 0 {
			length += v
		} else {
			length += v - 1
			hasOdd = true
		}
	}
	if hasOdd {
		length++
	}
	return length
}
