package arrayandmatrix

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

// 同余定理
// 对于(sum[j]−sum[i])%K==0 等于 sum[j]%K == sum[i]%K
// 对于nums[i:j]可被K整除的下标i，使用sum[i]%K作为map的key值，发现有与其相同的余数val++
// C(2,n) = n(n-1)/2，有n个相同的key，那么在统计的过程中可以通过++实现计数
// A B C 两两组合个数3，加入D后A B C D个数为6，3 + 3 =6
func subarraysDivByK1(A []int, K int) int {
	ret, sum := 0, 0
	m := make(map[int]int)
	m[0] = 1 // 对于[5,0]这样的子数组，计数应该是1+2，而不是1，特殊情况本身nums[i]%K==0，需要加上其自身
	for i := 0; i < len(A); i++ {
		sum += A[i]
		key := (sum%K + K) % K // 取模
		ret += m[key]
		m[key]++
	}
	return ret
}
