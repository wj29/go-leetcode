package arrayandmatrix

import (
	"testing"
)

// 给你一个长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积
// 请不要使用除法，且在 O(n) 时间复杂度内完成此题
// https://leetcode-cn.com/problems/product-of-array-except-self medium
func Test_ProductExceptSelf(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	t.Log(productExceptSelf(nums))
}

// 以nums[i]建立左右两侧乘积和，那么对于结果i即使l[i]*r[i]
func productExceptSelf(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	l, r, ret := make([]int, len(nums)), make([]int, len(nums)), make([]int, len(nums))
	l[0], r[len(nums)-1] = 1, 1 // 对于左右乘积和第一个元素没有前一个元素，那么它的值为0
	for i := 1; i < len(nums); i++ {
		l[i] = l[i-1] * nums[i-1]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		r[i] = r[i+1] * nums[i+1]
	}
	for i := range ret {
		ret[i] = l[i] * r[i]
	}
	return ret
}
