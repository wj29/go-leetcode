package dynamicprogramming

import (
	"testing"
	"github.com/wujiang1994/go-leetcode/common"
)

// 给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积
// https://leetcode-cn.com/problems/maximum-product-subarray/ medium
func Test_MaxProduct(t *testing.T) {
	//nums := []int{2, 3, -2, 4}
	nums := []int{-4, -3, -2}
	t.Log(maxProduct(nums))
}

// 此题应该是和最大的连续子数组的进阶版，详见/max_sub_array_test.go
// 由于乘法存在正负的区别，负负得正，要求得f(n)的最大值，我们应该知道f(n-1)的最大值和最小值
func maxProduct(nums []int) int {
	if len(nums) <= 0 {
		panic("invalid length")
	}
	max, min, ret := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		max, min = common.Max(nums[i], common.Max(max*nums[i], min*nums[i])), common.Min(nums[i], common.Min(max*nums[i], min*nums[i]))
		ret = common.Max(max, ret)
	}
	return ret
}
