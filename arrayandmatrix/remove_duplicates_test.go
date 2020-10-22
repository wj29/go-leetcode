package arrayandmatrix

import "testing"

// 80. 删除排序数组中的重复项 II
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii/
func Test_RemoveDuplicates(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	t.Log(removeDuplicates(nums))
}

// 此方法得到长度
func removeDuplicates(nums []int) int {
	begin, end := 0, 0
	for end < len(nums) {
		begin = end
		end++
		for end < len(nums) {
			if nums[end] != nums[end-1] {
				break
			}
			end++
		}
		if end-begin > 2 {
			nums = append(nums[:begin+2], nums[end:]...)
			end = begin + 2
		}
	}
	return len(nums)
}
