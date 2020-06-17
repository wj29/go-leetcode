package greedy

import "testing"

// 给你一个长度为 n 的整数数组，请你判断在 最多 改变 1 个元素的情况下，该数组能否变成一个非递减数列
// 非递减数列的：对于数组中所有的 i (1 <= i < n)，总满足 array[i] <= array[i + 1]
// https://leetcode-cn.com/problems/non-decreasing-array  easy
func Test_CheckPossibility(t *testing.T) {
	nums := []int{4, 2, 3}
	t.Log(checkPossibility(nums))
}

func checkPossibility(nums []int) bool {
	if len(nums) < 3 {
		return true
	}
	p := false
	// 第 0 个的处理比较特殊，因为它不需要照顾前面的数，所以直接改小自己就行
	if nums[0] > nums[1] {
		nums[0] = nums[1]
		p = true
	}
	for i := 1; i < len(nums)-1; i++ {
		// 当前数大于后一个数，破坏了非递减
		if nums[i] > nums[i+1]{
			if p {
				return false
			}
			p = true
			// 改动时，要 nums[i] 变小还是 nums[i+1] 变大需要做一个判断
			// nums[i] 和 nums[i-1] 的关系是确定的， nums[i] >= nums[i - 1]
			// nums[i] 也大于 nums[i + 1]，现在需要确定 nums[i + 1] 和 nums[i - 1] 的关系
			// 如果nums[ i + 1] 大于 nums[i - 1]，则调小 nums[i] 和 调大 nums[i + 1] 都行
			// 但是这里还是要优先调小 nums[i]，因为 nums[i + 1] 大了更不利于非减的要求，它可能大于后面的数了
			// 如果 nums[i + 1] 小于 nums[i - 1]，则必须调大 nums[i + 1]
			if nums[i+1] >= nums[i-1] {
				nums[i] = nums[i-1]
			} else {
				nums[i+1] = nums[i]
			}
		}
	}
	return true
}
