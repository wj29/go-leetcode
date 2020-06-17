package doublepoint

import (
	"sort"
	"testing"
)

// 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组
// 注意：答案中不可以包含重复的三元组
// https://leetcode-cn.com/problems/3sum
func Test_ThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	nums = []int{0, 0, 0}
	t.Log(threeSum(nums))
}

// 基础题twosum可以通过map实现N的时间复杂度找到目标target
// 按照此思想，我们知道a+b+c=target，同样可以找到a+b=target-c找到是否存在这个值是否在数组中，时间复杂度N^2
// 优化思路，通过将数组元素排序，跳过相同的元素达到去重的效果，如果nums[i]>target，i之后的元素不需要考虑
// 同时用l r双指针将三个数相加，比较与target的大小移动l和r
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	ret := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := len(nums) - 1
		for l < r {
			if nums[i]+nums[l]+nums[r] == 0 {
				ret = append(ret, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				l++
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				r--
			} else if nums[i]+nums[l]+nums[r] > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return ret
}
