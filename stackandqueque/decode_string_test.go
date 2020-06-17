package stackandqueque

import (
	"strconv"
	"testing"
)

func Test_DecodeString(t *testing.T) {
	s := "3[a]2[bc]"
	s = "3[a2[c]]"
	s = "11[a]"
	t.Log(decodeString(s))
}

// 通过栈实现，遇到第一个']'前一直入栈，遇到']'后开始出栈，直到遇到第一个'['为止，下一个出栈的必定是数字
// 得到出栈的字符串后翻转后进行解码，将解码后的字符串重新入栈直到字符串被遍历完
// 数字和'['一定是一起出现
func decodeString(s string) string {
	ret := ""
	stack := make([]string, 0)
	for i := 0; i < len(s); i++ {
		if string(s[i]) != "]" { // 入栈
			stack = append(stack, string(s[i]))
		} else {
			tmp := make([]string, 0)
			for j := len(stack) - 1; j >= 0; j-- {
				if stack[j] == "[" { // 出栈到第一个'['
					stack = stack[:j]
					break
				} else {
					tmp = append(tmp, stack[j])
				}
			}
			for k := 0; k < len(tmp)/2; k++ { // 翻转
				tmp[k], tmp[len(tmp)-1-k] = tmp[len(tmp)-1-k], tmp[k]
			}
			index := 0
			for l := len(stack) - 1; l >= -1; l-- { // 可能是多位数
				if l == -1 { // 第零位特殊处理
					index = l
					break
				}
				if stack[l] < "0" || stack[l] > "9" {
					index = l
					break
				}
			}
			str := ""
			for n := index + 1; n < len(stack); n++ {
				str += stack[n]
			}
			co, _ := strconv.Atoi(str)
			stack = stack[:index+1]
			for p := 0; p < co; p++ {
				stack = append(stack, tmp...)
			}
		}
	}
	for i := 0; i < len(stack); i++ {
		ret += stack[i]
	}
	return ret
}
