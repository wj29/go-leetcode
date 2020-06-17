package dynamicprogramming

import (
	"testing"
)

// 给定一个整数数组 A，返回其中元素之和可被 K 整除的（连续、非空）子数组的数目
// https://leetcode-cn.com/problems/subarray-sums-divisible-by-k/ medium
func Test_SubarraysDivByK(t *testing.T) {
	//A := []int{4, 5, 0, -2, -3, 1}
	//K := 5
	A := []int{-1, 2, 9}
	K := 2
	t.Log(subarraysDivByK(A, K))
	t.Log(subarraysDivByK1(A, K))
}

// 连续子数组和，可以通过构造前缀和来实现
// 构造前缀和，循环数组即可知道所有的连续子数组和
// 时间复杂度为N^2，leetcode上超出时间限制
func subarraysDivByK(A []int, K int) int {
	if len(A) == 0 {
		return 0
	}
	prev := make([]int, len(A)+1)
	for i := 0; i < len(A); i++ {
		if i == 0 {
			prev[i+1] = A[i]
		} else {
			prev[i+1] = prev[i] + A[i]
		}
	}
	count := 0
	for i := 1; i < len(prev); i++ {
		for j := 0; j < i; j++ {
			if (prev[i]-prev[j])%K == 0 { // 同余定理
				count++
			}
		}
	}
	return count
}

// (prev[i]-prev[j])%K == 0 同余定理 可以表示成 prev[i]modK == prev[j]modK，那么只需要一次遍历即可
// 排列组合中两两组合的个数 + 加上一个数据前的长度 = 变换后两两组合的个数
// A B C 两两组合个数3，加入D后A B C D个数为6，3 + 3 =6
// 即Cn 2 + n = Cn+1 2
func subarraysDivByK1(A []int, K int) int {
	record := map[int]int{0: 0}
	sum, ans := 0, 0
	B := make([]int, 0)
	B = append(B, 0)
	B = append(B, A...)
	for _, elem := range B {
		sum += elem
		modulus := (sum % K + K) % K
		//modulus := sum % K // 错误
		ans += record[modulus]
		record[modulus]++
	}
	return ans
}
