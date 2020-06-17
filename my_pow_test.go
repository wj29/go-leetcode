package leetcode

import (
	"testing"
)

// 实现pow函数，即x的y次方
// https://leetcode-cn.com/problems/powx-n/ medium
func Test_MyPow(t *testing.T) {
	t.Log(myPow(2, 4))
}

// x的y次方即x的y/2次方的平方，迭代即可，时间复杂度log2(N)
func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	}
	return 1 / quickMul(x, -n)
}

func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := quickMul(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}
