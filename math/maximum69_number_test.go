package math

import (
	"math"
	"strconv"
	"testing"
)

// 给你一个仅由数字 6 和 9 组成的正整数 num。
// 你最多只能翻转一位数字，将 6 变成 9，或者把 9 变成 6 。
// 请返回你可以得到的最大数字。
// https://leetcode-cn.com/problems/maximum-69-number/
func Test_Maximum69Number(t *testing.T) {
	num := 66999
	t.Log(maximum69Number(num))
}

// 从头往后遍历找到第一个6即可
func maximum69Number(num int) int {
	str := strconv.Itoa(num)
	index := -1
	for i := 0; i < len(str); i++ {
		if str[i] == '6' {
			index = i
			break
		}
	}
	if index > -1 {
		num += int(3 * math.Pow10(len(str)-index-1))
	}
	return num
}
