package doublepoint

import "testing"

// 33. 搜索旋转排序数组
// https://leetcode-cn.com/problems/search-in-rotated-sorted-array/
func Test_Search(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	t.Log(search(nums, target))
}

// 将数组一分为二，其中一定有一个是有序的，另一个是部分有序，此时有序部分用二分法查找，无序的循环调用本身
func search(nums []int, target int) int {
	return search2(nums, 0, len(nums)-1, target)
}

func search2(nums []int, start, end, target int) int {
	if start > end {
		return -1
	}
	l, r := start, end

	mid := l + (r-l)>>1
	if nums[mid] >= nums[l] {
		if v := find(nums, l, mid, target); v >= 0 {
			return v
		} else {
			return search2(nums, mid+1, r, target)
		}
	} else {
		if v := find(nums, mid+1, r, target); v >= 0 {
			return v
		} else {
			return search2(nums, l, mid, target)
		}
	}
}

func find(nums []int, start, end, target int) int {
	l, r := start, end
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		}
	}
	return -1
}
