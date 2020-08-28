package greedy

import (
	"strings"
	"testing"
)

// 12. 整数转罗马数字
// https://leetcode-cn.com/problems/integer-to-roman/
func Test_IntToRoman(t *testing.T) {
	num := 53
	t.Log(intToRoman(num))
}

// 罗马数字转换规则
// 从左到右选择尽可能大的符号表示
// 此时即可选择贪心处理，优先选择最大的符号表示
func intToRoman(num int) string {
	values := [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	resp := [13]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	ret := ""
loop:
	for num != 0 {
		for i := 0; i < len(values); i++ {
			if num >= values[i] {
				num -= values[i]
				ret += resp[i]
				continue loop
			}
		}
	}
	return ret
}

// 13. 罗马数字转整数
// https://leetcode-cn.com/problems/roman-to-integer/
func romanToInt(s string) int {
	values := [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	resp := [13]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	ret := 0
loop:
	for len(s) != 0 {
		for i := 0; i < len(resp); i++ {
			if strings.HasPrefix(s, resp[i]) {
				s = s[len(resp[i]):]
				ret += values[i]
				continue loop
			}
		}
	}
	return ret
}
