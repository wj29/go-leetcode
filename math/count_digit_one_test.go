package math

import (
	"strconv"
	"strings"
	"testing"
)

// 给定一个整数 n，计算所有小于等于 n 的非负整数中数字 1 出现的个数
// https://leetcode-cn.com/problems/number-of-digit-one/ hard
func Test_CountDigitOne(t *testing.T) {
	n := 824883294
	t.Log(countDigitOne(n))
}

// 超时
func countDigitOne(n int) int {
	ret := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if i == 1 {
			ret[i] = 1
		} else {
			ret[i] = ret[i-1] + strings.Count(strconv.Itoa(i), "1")
		}
	}
	return ret[n]
}
