package dynamicprogramming

import "testing"

// 数组 A 包含 N 个数，且索引从0开始。数组 A 的一个子数组划分为数组 (P, Q)，P 与 Q 是整数且满足 0<=P<Q<N
// 如果满足以下条件，则称子数组(P, Q)为等差数组：
// 元素 A[P], A[p + 1], ..., A[Q - 1], A[Q] 是等差的。并且 P + 1 < Q 。
// 函数要返回数组 A 中所有为等差数组的子数组个数
// https://leetcode-cn.com/problems/arithmetic-slices  medium
func Test_NumberOfArithmeticSlices(t *testing.T) {
	A := []int{1, 2, 3, 4}
	t.Log(numberOfArithmeticSlices(A))
}

// 对于dp[i][j](j-i>1)来说，如果是等差数列，那么dp[i-1][j]为等差数列的条件为nums[i][j]-nums[i-1][j]的差也等于该等差数列的差
// 对于dp[i][j+1]同理
// base: dp[i][j](j-i<2)都不是等差数列，计算出所有dp[i][i+2]是否为等差数列
// 同 longest_palindrome_test.go
func numberOfArithmeticSlices(A []int) int {
	if len(A) <= 2 {
		return 0
	}
	ret := 0
	dp := make([][]bool, len(A))
	for i := range dp {
		dp[i] = make([]bool, len(A))
	}
	for gap := 0; gap < len(A); gap++ {
		for i := 0; i+gap < len(A); i++ {
			j := i + gap
			if gap <= 1 {
				dp[i][j] = false
			} else if gap == 2 {
				dp[i][j] = A[i+2]-A[i+1] == A[i+1]-A[i]
				if dp[i][j] {
					ret++
				}
			} else {
				dp[i][j] = (dp[i+1][j] && A[i+2]-A[i+1] == A[i+1]-A[i]) || (dp[i][j-1] && A[j]-A[j-1] == A[j-1]-A[j-2])
				if dp[i][j] {
					ret++
				}
			}
		}
	}
	return ret
}
