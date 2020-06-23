package arrayandmatrix

import (
	"fmt"
	"testing"
)

// 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
//( 例如，数组 [0,0,1,2,2,5,6] 可能变为 [2,5,6,0,0,1,2] )。
// 编写一个函数来判断给定的目标值是否存在于数组中。若存在返回 true，否则返回 false。
// https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii medium
func Test_Search(t *testing.T) {
	nums := []int{3, 1}
	target := 1
	t.Log(search(nums, target))
}

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	index := len(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			index = i + 1
			break
		}
	}
	if target == nums[0] {
		return true
	} else if target > nums[0] {
		return dichotomy(nums[:index], target)
	} else {
		return dichotomy(nums[index:], target)
	}
}

func dichotomy(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			return true
		} else if nums[mid] > target {
			r = mid
		} else {
			l = mid + 1
		}
		if l == r {
			return nums[l] == target
		}
	}
	return false
}

func removeFromzSet(zSet []string, value string) []string {
	for i, v := range zSet {
		if v == value {
			zSet = append(zSet[:i], zSet[i+1:]...)
			fmt.Println(zSet)
			return zSet
		}
	}
	return []string{}
}
