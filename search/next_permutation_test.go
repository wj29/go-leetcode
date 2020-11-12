package search

import (
    "sort"
    "testing"
)

// 31. 下一个排列
// https://leetcode-cn.com/problems/next-permutation/
func Test_NextPermutation(t *testing.T) {
    nums := []int{1, 2, 3}
    nextPermutation(nums)
    t.Log(nums)
}

// 下一个更大的排列，高位后动，即需要从后向前遍历，当前一个数字依次比后一个数字大时即不存在下一个更大的排列
// 当出现一个数字比后一个数字小时，那么就会有下一个更大的排列，包含该位的数已经上最大的排列，需要找到第一个比该数字大的数字放到该位，其他数字从小到大顺序排列即可
func nextPermutation(nums []int) {
    for i := len(nums) - 1; i >= 0; i-- {
        if i != 0 && nums[i] <= nums[i-1] {
            continue
        }
        index := i - 1
        if index < 0 {
            for j := 0; j < len(nums)/2; j++ {
                nums[j], nums[len(nums)-j-1] = nums[len(nums)-j-1], nums[j]
            }
            return
        }
        for k := len(nums) - 1; k > index; k-- {
            if nums[k] <= nums[index] {
                continue
            }
            nums[k], nums[index] = nums[index], nums[k]
            sort.Ints(nums[index+1:])
            return
        }
    }
}
