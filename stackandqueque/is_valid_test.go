package stackandqueque

import "testing"

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
// 有效字符串需满足：
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 注意空字符串可被认为是有效字符串。
// https://leetcode-cn.com/problems/valid-parentheses

func Test_IsValid(t *testing.T) {
	s := "()[]{}"
	t.Log(isValid(s))
}

// 遇到括号左边入栈，右边出栈
func isValid(s string) bool {
	stack := make([]byte, 0)
	m := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	for len(s) != 0 {
		current := s[0]
		s = s[1:]
		if current == '(' || current == '{' || current == '[' {
			stack = append(stack, current)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != m[current] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return len(s) == 0 && len(stack) == 0
}
