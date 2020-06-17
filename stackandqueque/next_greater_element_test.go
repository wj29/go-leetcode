package stackandqueque

import "testing"

// 给定两个 没有重复元素 的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。找到 nums1 中每个元素在 nums2 中的下一个比其大的值。
// nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出 -1
// https://leetcode-cn.com/problems/next-greater-element-i  easy

// 给定一个循环数组（最后一个元素的下一个元素是数组的第一个元素），输出每个元素的下一个更大元素。数字 x 的下一个更大的元素是按数组
// 遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1
// https://leetcode-cn.com/problems/next-greater-element-ii medium
func Test_NextGreaterElement(t *testing.T) {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	t.Log(nextGreaterElement1(nums1, nums2))

	nums := []int{1, 2, 1}
	t.Log(nextGreaterElements(nums))
}

// 对于nums1中每个元素，遍历nums2中的没元素，找到第一个比他大的元素输出
// 时间复杂度len(nums1)*len(nums2)
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ret := make([]int, len(nums1))
	flag := false
	for i := 0; i < len(nums1); i++ {
		ret[i] = -1
		flag = false
		for j := 0; j < len(nums2); j++ {
			if nums2[j] == nums1[i] {
				flag = true
			}
			if flag && nums2[j] > nums1[i] {
				ret[i] = nums2[j]
				break
			}
		}
	}
	return ret
}

// 通过单调栈实现
// 对于数组nums2的每一个元素，从头开始入栈，当nums[i+1]<=栈顶时或者栈为空时将nums[i+1]进行入栈，此时构建了一个从顶到下的单调不递减的栈
// 当nums[i+1] > 栈顶元素时找到了栈中所有元素的第一个大的元素，构建map存储，对所有栈元素进行出栈
// 时间复杂度N
func nextGreaterElement1(nums1 []int, nums2 []int) []int {
	stack := make([]int, 0)
	m := make(map[int]int)
	ret := make([]int, len(nums1))
	for i := 0; i < len(nums2); i++ {
		if len(stack) == 0 { // 空栈直接入栈
			stack = append(stack, nums2[i])
			continue
		}
		if nums2[i] <= stack[len(stack)-1] { // 比栈顶小入栈
			stack = append(stack, nums2[i])
		} else {
			for len(stack) != 0 { // 逐一比较栈顶
				index := len(stack) - 1
				if nums2[i] <= stack[index] {
					break
				}
				m[stack[index]] = nums2[i]
				stack = stack[:index] // 出栈
			}
			stack = append(stack, nums2[i]) // 新元素入栈
		}
	}
	for i := range stack {
		m[stack[i]] = -1
	}
	for i := range nums1 {
		ret[i] = m[nums1[i]]
	}
	return ret
}

// 题2中没有注明数组中元素不能相等，当然也没有nums1需要去用map匹配，我们可以在循环或者构建栈的时候直接对结果进行记录
// 循环数组，即将数组重新append到它本身后面即可以构成循环数组
// 在栈中不能通过数组元素做匹配，不能保证顺序，栈中元素应该是数组索引
func nextGreaterElements(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	ret := make([]int, 2*len(nums)-1)
	for i := range ret {
		ret[i] = -1
	}
	stack := make([]int, 2*len(nums)-1)

	nums = append(nums, nums...)
	for i := 0; i < len(nums); i++ {
		if len(stack) == 0 {
			stack = append(stack, i)
			continue
		}
		if nums[i] <= nums[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else {
			for len(stack) != 0 {
				index := len(stack) - 1
				if nums[i] <= nums[stack[index]] {
					break
				}
				ret[stack[index]] = nums[i]
				stack = stack[:index]
			}
			stack = append(stack, i)
		}
	}
	return ret[:(len(nums)+1)/2]
}
