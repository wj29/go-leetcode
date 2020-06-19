package math

import (
	"strconv"
	"testing"
)

// TODO
// 给定一个非负整数 num，反复将各个位上的数字相加，直到结果为一位数
// 输入: 38 输出: 2 各位相加的过程为：3 + 8 = 11, 1 + 1 = 2
// https://leetcode-cn.com/problems/add-digits/  easy
// 进阶: 1的时间复杂度
func Test_AddDigits(t *testing.T) {
	num := 38
	t.Log(addDigits1(num))
}

func addDigits(num int) int {
	numStr := strconv.Itoa(num)
	if len(numStr) < 2 {
		return num
	}
	for len(numStr) > 1 {
		tmp := 0
		for i := 0; i < len(numStr); i++ {
			t, _ := strconv.Atoi(string(numStr[i]))
			tmp += t
		}
		numStr = strconv.Itoa(tmp)
	}
	num, _ = strconv.Atoi(numStr)
	return num
}

// 没看懂解法
func addDigits1(num int) int {
	return (num-1)%9 + 1
}
