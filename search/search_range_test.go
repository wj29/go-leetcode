package search

import (
	"testing"
)

// 34. 在排序数组中查找元素的第一个和最后一个位置
// https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
// 时间复杂度 O(log n)
func Test_SearchRange(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	nums = []int{1}
	target = 1
	t.Log(searchRange(nums, target))
}

// 二分法的思路
// 注意 for left < right 表示，当left == right的时候就停止查找，此时可以检查这个元素是不是我们要找的
// for left <= right 表示 left == right 的时候还继续查找，left == right+1时停止
func searchRange(nums []int, target int) []int {
	ret := []int{-1, -1}
	if len(nums) == 0 || nums[0] > target || nums[len(nums)-1] < target {
		return ret
	}
	l, r := 0, len(nums)
	for l < r {
		mid := l + (r-l)>>1
		if nums[mid] == target { // 不停止，继续向左侧搜索
			r = mid
		} else if nums[mid] > target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	// 返回r
	if nums[r] != target {
		return ret
	}
	ret[0] = r

	r = len(nums)
	for l < r {
		mid := l + (r-l)>>1
		if nums[mid] == target { // 继续向右侧搜索
			l = mid + 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	ret[1] = l - 1

	return ret
}
