package math

import (
	"strconv"
	"testing"
)

// 给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
// https://leetcode-cn.com/problems/multiply-strings/
func Test_Multiply(t *testing.T) {
	num1, num2 := "123", "456"
	t.Log(multiply(num1, num2))
}

func Test_multiplySingle(t *testing.T) {
	r := multiplySingle("789", 9, 2)
	t.Log(789 * 900)
	t.Log(r)
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	if len(num1) < len(num2) { // keep num2 length shorter
		num1, num2 = num2, num1
	}
	ret := ""
	for i := len(num2) - 1; i >= 0; i-- { // calculate from end to start
		ret = add2(ret, multiplySingle(num1, int(num2[i]-'0'), len(num2)-1-i))
	}
	return ret
}

// multi-digit and single-digit multi, cnt means the location of single-digit in original string
func multiplySingle(s string, m, cnt int) string {
	carry, ret := 0, ""
	for i := len(s) - 1; i >= 0; i-- {
		val := int(s[i]-'0')*m + carry // calculate
		carry = val / 10
		val %= 10
		ret = strconv.Itoa(val) + ret
	}
	if carry > 0 {
		ret = strconv.Itoa(carry) + ret // the last carry
	}
	for i := 1; i <= cnt; i++ {
		ret += "0"
	}
	return ret
}

func add2(s1 string, s2 string) string {
	if len(s1) < len(s2) { // keep s2 length shorter
		s1, s2 = s2, s1
	}
	sub := len(s1)-len(s2)
	for i := 1; i <= sub; i++ {
		s2 = "0" + s2
	}
	carry, ret := 0, ""
	for i := len(s2) - 1; i >= 0; i-- {
		val := int(s1[i]-'0') + int(s2[i]-'0') + carry
		carry = val / 10
		val %= 10
		ret = strconv.Itoa(val) + ret
	}
	if carry > 0 {
		ret = strconv.Itoa(carry) + ret // the last carry
	}
	return ret
}
