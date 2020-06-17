package doublepoint

import (
	"math"
	"sort"
	"testing"
)

// 给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案
// https://leetcode-cn.com/problems/3sum-closest  medium
func Test_ThreeSumClosest(t *testing.T) {
	nums := []int{1, 1, 1, 1}
	target := 0
	t.Log(threeSumClosest(nums, target))
}

// 和三数之和一样，将差值最小的存下来，逐一比较即可
func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	ret := make([]int, 2)
	ret[0] = math.MaxInt32
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := len(nums) - 1

		for l < r {
			if nums[i]+nums[l]+nums[r] == target {
				return target
			} else if nums[i]+nums[l]+nums[r] > target {
				if int(math.Abs(float64(nums[i]+nums[l]+nums[r]-target))) < ret[0] {
					ret[0] = int(math.Abs(float64(nums[i] + nums[l] + nums[r] - target)))
					ret[1] = nums[i] + nums[l] + nums[r]
				}
				r--
			} else {
				if int(math.Abs(float64(nums[i]+nums[l]+nums[r]-target))) < ret[0] {
					ret[0] = int(math.Abs(float64(nums[i] + nums[l] + nums[r] - target)))
					ret[1] = nums[i] + nums[l] + nums[r]
				}
				l++
			}
		}
	}
	return ret[1]
}
