package arrayandmatrix

import "testing"

// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序
// 必须在原数组上操作，不能拷贝额外的数组
// https://leetcode-cn.com/problems/move-zeroes/description/  easy
func Test_MoveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	t.Log(nums)
}

// 双指针遍历，遇到0的则与下一位交换
func moveZeroes(nums []int) {
	i, j := 0, 0
	for j < len(nums) {
		if i != j {
			nums[i], nums[j] = nums[j], nums[i]
		}
		if nums[i] != 0 {
			i++
		}
		j++
	}
}
