package dynamicprogramming

import "testing"

// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢
// https://leetcode-cn.com/problems/climbing-stairs/description/  easy
func Test_ClimbStairs(t *testing.T) {
	n := 10
	t.Log(climbStairs(n))
}

// f(n) = f(n-1) + f(n-2)
func climbStairs(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	ret := make([]int, n)
	ret[0] = 1
	ret[1] = 2
	for i := 2; i < n; i++ {
		ret[i] = ret[i-1] + ret[i-2]
	}
	return ret[n-1]
}
