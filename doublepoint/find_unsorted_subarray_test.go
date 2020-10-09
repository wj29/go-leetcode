package doublepoint

import (
	"testing"
)

// 581. 最短无序连续子数组
// https://leetcode-cn.com/problems/shortest-unsorted-continuous-subarray/
func Test_FindUnsortedSubarray(t *testing.T) {
	nums := []int{2, 6, 4, 8, 10, 9, 15}
	t.Log(findUnsortedSubarray(nums))
}

// 连续子数组，先找到左右端已经排好序的最长两端，得到最短的中间乱序连续子数组
// 获取其上线边界，判断在左右两端的位置即可
func findUnsortedSubarray(nums []int) int {
	l, r := 0, len(nums)-1
	for l < len(nums)-1 && nums[l] <= nums[l+1] { // 排好序的左边界
		l++
	}
	for r > 0 && nums[r] >= nums[r-1] { // 排好序的有右边界
		r--
	}
	if l >= r { // 表明是顺序的
		return 0
	}
	down, up := getBorder(nums, l, r) // 获取乱序子数组上下边界
	// 在排好序的左右数组中，对l和r的区域进行扩大
	for l > 0 {
		if nums[l-1] > down {
			l--
		} else {
			break
		}
	}
	for r < len(nums)-1 {
		if nums[r+1] < up {
			r++
		} else {
			break
		}
	}
	return r - l + 1
}

func getBorder(nums []int, l, r int) (int, int) {
	max, min := nums[l], nums[l]
	for i := l; i <= r; i++ {
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min, max
}
