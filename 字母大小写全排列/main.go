package main

// 给定一个字符串 s ，通过将字符串 s 中的每个字母转变大小写，我们可以获得一个新的字符串。
// 返回所有可能得到的字符串集合 。以任意顺序返回输出。

func letterCasePermutation(s string) []string {
	result := []string{}
	var dfs func(index int, path []byte)
	dfs = func(index int, path []byte) {
		if index == len(s) {
			result = append(result, string(path))
			return
		}
		ch := s[index]
		if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' {
			dfs(index+1, append(path, ch))
			if ch >= 'a' && ch <= 'z' {
				dfs(index+1, append(path, ch-'a'+'A'))
			} else {
				dfs(index+1, append(path, ch-'A'+'a'))
			}
		} else {
			dfs(index+1, append(path, ch))
		}
	}
	dfs(0, []byte{})
	return result
}
