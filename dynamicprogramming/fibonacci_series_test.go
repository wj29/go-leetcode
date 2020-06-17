package dynamicprogramming

import "testing"

// 斐波那契数列
// f(0)=0 f(1)=1 f(n)=f(n-1)+f(n-2)
// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233
func Test_Fibonacci_Series(t *testing.T) {
	t.Log(fibonacciSeries(100))
}

func fibonacciSeries(n int) int {
	ret := make([]int, n+1) // 构建结果储存加以复用
	for i := 0; i < n+1; i++ {
		if i < 2 {
			ret[i] = i
		} else {
			ret[i] = ret[i-1] + ret[i-2]
		}
	}
	return ret[n]
}
