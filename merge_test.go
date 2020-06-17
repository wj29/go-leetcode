package leetcode

import (
	"testing"
)

// 两个有序数组m和n，归并两个有序数组成为一个有序数组 easy
// https://leetcode-cn.com/problems/merge-sorted-array/description/

func Test_Merge(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{1, 2, 3}
	m := 6
	n := 3
	merge(nums1, m, nums2, n)
	t.Log(nums1)
}

// 插入排序，或者归并排序都行  // 归并排序内存消耗很大
func merge(nums1 []int, m int, nums2 []int, n int) {
	var leftIndex, rightIndex int
	var sortedPile []int
	for leftIndex < m && rightIndex < n {
		if nums1[leftIndex] < nums2[rightIndex] {
			sortedPile = append(sortedPile, nums1[leftIndex])
			leftIndex++
		} else if nums1[leftIndex] > nums2[rightIndex] {
			sortedPile = append(sortedPile, nums2[rightIndex])
			rightIndex++
		} else {
			sortedPile = append(sortedPile, nums1[leftIndex])
			leftIndex++
			sortedPile = append(sortedPile, nums2[rightIndex])
			rightIndex++
		}
	}

	for leftIndex < m {
		sortedPile = append(sortedPile, nums1[leftIndex])
		leftIndex++
	}
	for rightIndex < n {
		sortedPile = append(sortedPile, nums2[rightIndex])
		rightIndex++
	}
	copy(nums1, sortedPile)
}
