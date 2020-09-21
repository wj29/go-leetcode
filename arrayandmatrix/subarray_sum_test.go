package arrayandmatrix

import (
	"testing"
)

// 给定一个整数数组和一个整数k，你需要找到该数组中和为k的连续的子数组的个数
// 输入:nums = [1,1,1], k = 2
// 输出: 2 , [1,1] 与 [1,1] 为两种不同的情况
// https://leetcode-cn.com/problems/subarray-sum-equals-k/  medium
func Test_SubarraySum(t *testing.T) {
	nums := []int{1, 1, 1}
	t.Log(subarraySum(nums, 2))
}

// 前缀和
// 通过构造前缀和，可以在N^2的时间复杂度得到所有和为K的数组
func subarraySum(nums []int, k int) int {
	pre := make([]int, len(nums)+1)
	pre[0] = 0
	for i := 0; i < len(nums); i++ {
		pre[i+1] += nums[i] + pre[i]
	}
	count := 0
	for i := 1; i < len(pre); i++ {
		for j := 0; j < i; j++ {
			if pre[i]-pre[j] == k {
				count++
			}
		}
	}
	return count
}

// 对于和为K的子数组，nums[i:j]即满足prev[j] - prev[i] == K
// 即prev[i] == prev[j] - K
// 将prev[i]作为map的key，向后统计，满足prev[j]-K == key时，val++
// 统计出共有多少个何为K的子数组
func subarraySum2(nums []int, k int) int {
	ret, sum := 0, 0
	m := make(map[int]int)
	m[0] = 1 // 对于[5,0]这样的子数组，计数应该是1+1，而不是1，特殊情况本身nums[i] == k，需要加上其自身
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if v, ok := m[sum-k]; ok {
			ret += v
		}
		m[sum]++
	}
	return ret
}
