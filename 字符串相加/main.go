package main

import (
	"strconv"
	"strings"
)

// 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和并同样以字符串形式返回。
// 你不能使用任何內建的用于处理大整数的库（比如 BigInteger）， 也不能直接将输入的字符串转换为整数形式。

func addStrings(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}
	num2 = strings.Repeat("0", len(num1)-len(num2)) + num2
	num := 0
	result := []int{}
	for i := len(num1) - 1; i >= 0; i-- {
		a, _ := strconv.Atoi(string(num1[i]))
		b, _ := strconv.Atoi(string(num2[i]))
		sum := a + b + num
		num = sum / 10
		result = append(result, sum%10)
	}
	if num > 0 {
		result = append(result, num)
	}
	var s strings.Builder
	for i := len(result) - 1; i >= 0; i-- {
		s.WriteString(strconv.Itoa(result[i]))
	}
	return s.String()
}
