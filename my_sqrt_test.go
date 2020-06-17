package leetcode

import "testing"

// 实现 int sqrt(int x) 函数
// 计算并返回 x 的平方根，其中 x 是非负整数
// 由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去
// https://leetcode-cn.com/problems/sqrtx/ easy
func Test_MySqrt(t *testing.T) {
	t.Log(mySqrt(8))
}

// 二分法
func mySqrt(x int) int {
	l, r := 0, x
	ret := -1
	for l <= r {
		mid := (l + r) / 2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			ret = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ret
}
