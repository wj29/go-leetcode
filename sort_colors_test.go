package leetcode

import "testing"

// 荷兰国旗问题
// 给定一个包含红白蓝三种颜色共n个元素的数组，原地进行排序，使得红白蓝排序并颜色相同的相邻 medium
// 此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色
// https://leetcode-cn.com/problems/sort-colors/

func Test_SortColors(t *testing.T) {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	t.Log(nums)
}

// 使用left right current 三个指针进行遍历，寻找红蓝的边界并交换 时间负载度N 空间1
func sortColors(nums []int) {
	left := 0
	right := len(nums) - 1
	current := 0

	for {
		if nums[current] == 0 {
			swap(nums, left, current)
			left++
		} else if nums[current] == 1 {
			current++
		} else {
			swap(nums, current, right)
			right--
		}
		if current < left {
			current = left
		}
		if current > right {
			break
		}
	}
}
