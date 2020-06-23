package math

import (
	"testing"
)

// 给你两个二进制字符串，返回它们的和（用二进制表示）。
// 输入为 非空 字符串且只包含数字 1 和 0。
// https://leetcode-cn.com/problems/add-binary/ easy
func Test_AddBinary(t *testing.T) {
	a := "1010"
	b := "1011"
	t.Log(addBinary(a, b))
}

// 不可以转换成二进制数字进行运算，超出长度
// 从尾部遍历向前进位即可，较短的字符前面补0
func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	var (
		c     string // 构造相同长度的字符
		carry bool   // 进位符
		value string // 计算后当前位的值
		ret   string // 结果
	)
	diff := len(a) - len(b)
	for diff > 0 {
		c += "0"
		diff--
	}
	c += b
	tmp := make([]string, 0)
	for i := len(a) - 1; i >= 0; i-- {
		value, carry = add(a[i], c[i], carry)
		tmp = append(tmp, value)
	}
	if carry {
		tmp = append(tmp, "1")
	}
	for i := len(tmp) - 1; i >= 0; i-- {
		ret += tmp[i]
	}
	return ret
}

func add(p1, p2 uint8, carry bool) (ret string, flag bool) {
	switch {
	case p1 == '0' && p2 == '0':
		ret = "0"
		if carry {
			ret = "1"
		}
	case p1 == '1' && p2 == '1':
		ret = "0"
		flag = true
		if carry {
			ret = "1"
		}
	case (p1 == '0' && p2 == '1') || (p1 == '1' && p2 == '0'):
		ret = "1"
		if carry {
			ret = "0"
			flag = true
		}
	}
	return
}
