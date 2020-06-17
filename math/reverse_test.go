package math

import (
	"math"
	"strconv"
	"testing"
)

// 给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转
// https://leetcode-cn.com/problems/reverse-integer/ easy
func Test_Reverse(t *testing.T) {
	x := 123
	t.Log(reverse(x))
}

// 数字转换成字符串反转即可
func reverse(x int) int {
	s := strconv.Itoa(x)
	if len(s) <= 1 {
		return x
	}
	ret := 0
	flag := false
	if string(s[0]) == "-" {
		flag = true
	}
	if flag {
		s = s[1:]
	}
	for i := len(s) - 1; i >= 0; i-- {
		n, _ := strconv.Atoi(string(s[i]))
		ret += n * int(math.Pow10(i))
	}
	if flag {
		ret = 0 - ret
	}
	if ret > math.MaxInt32 || ret < math.MinInt32 {
		return 0
	}
	return ret
}
