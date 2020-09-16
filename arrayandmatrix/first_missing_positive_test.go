package arrayandmatrix

import (
	"testing"
)

// 41. 缺失的第一个正数
// https://leetcode-cn.com/problems/first-missing-positive/
// 时间复杂度O(n) 常数空间复杂度
func Test_FirstMissingPositive(t *testing.T) {
	nums := []int{3, 4, 0, 1}
	nums = []int{2}
	t.Log(firstMissingPositive(nums))
}

// 利用另一个数组的形式可以直接做到，但是由于常数空间限制，利用数组本身的下标特性
// 将对应的数据放在对应的下标-1中(第i个元素的下标为i-1)
// 完成后循环数组判断当前下标内值是否为下标+1
func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 && nums[i] <= len(nums) && nums[i]-1 != i {
			if nums[nums[i]-1]-1 == nums[i]-1 { // 出现重复数据时，直接跳过，即需要交换的位置已经是正确的
				continue
			}
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			i-- // 回到上一个i
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}
