package main

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
// 有效字符串需满足：
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。

func isValid(s string) bool {
	// temp := make([]rune, 0)
	temp := []rune{}
	mapping := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	// for _, char := range s {
	// 	if char == '(' || char == '{' || char == '[' {
	// 		temp = append(temp, char)
	// 	} else if len(temp) == 0 {
	// 		return false
	// 	} else {
	// 		top := temp[len(temp)-1]
	// 		temp = temp[:len(temp)-1]

	// 		if top != mapping[char] {
	// 			return false
	// 		}
	// 	}
	// }

	for _, char := range s {
		if value, ok := mapping[char]; ok {
			if len(temp) == 0 || temp[len(temp)-1] != value {
				return false
			}
			temp = temp[:len(temp)-1]
		} else {
			temp = append(temp, char)
		}
	}

	return len(temp) == 0
}
