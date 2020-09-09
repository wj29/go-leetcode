package backtracking

import (
	"testing"
)

func Test_CanJump(t *testing.T) {
	nums := []int{3, 2, 1, 0, 4}
	nums = []int{2, 0, 6, 9, 8, 4, 5, 0, 8, 9, 1, 2, 9, 6, 8, 8, 0, 6, 3, 1, 2, 2, 1, 2, 6, 5, 3, 1, 2, 2, 6, 4, 2, 4, 3, 0, 0, 0, 3, 8, 2, 4, 0, 1, 2, 0, 1, 4, 6, 5, 8, 0, 7, 9, 3, 4, 6, 6, 5, 8, 9, 3, 4, 3, 7, 0, 4, 9, 0, 9, 8, 4, 3, 0, 7, 7, 1, 9, 1, 9, 4, 9, 0, 1, 9, 5, 7, 7, 1, 5, 8, 2, 8, 2, 6, 8, 2, 2, 7, 5, 1, 7, 9, 6}
	//t.Log(canJump(nums))
	nums = []int{2, 0}
	nums = []int{8, 2, 4, 4, 4, 9, 5, 2, 5, 8, 8, 0, 8, 6, 9, 1, 1, 6, 3, 5, 1, 2, 6, 6, 0, 4, 8, 6, 0, 3, 2, 8, 7, 6, 5, 1, 7, 0, 3, 4, 8, 3, 5, 9, 0, 4, 0, 1, 0, 5, 9, 2, 0, 7, 0, 2, 1, 0, 8, 2, 5, 1, 2, 3, 9, 7, 4, 7, 0, 0, 1, 8, 5, 6, 7, 5, 1, 9, 9, 3, 5, 0, 7, 5}
	t.Log(canJump2(nums))
}

// 超时
func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	m := make(map[int]bool)
	ret := false
	backTraceForCanJump(&ret, nums, 0, m)
	return ret
}

func backTraceForCanJump(ret *bool, nums []int, index int, m map[int]bool) {
	if *ret {
		return
	}
	if nums[index]+index >= len(nums)-1 {
		*ret = true
		return
	}
	for i := 0; i <= nums[index]; i++ {
		if m[index+i] {
			continue
		}
		m[index+i] = true
		backTraceForCanJump(ret, nums, index+i, m)
		m[index+i] = false
	}
}

// 是否能达到末尾，只要将能到达末尾的index筛选出来即可
// 因为只能向后跳，所以筛选的越来越少，使用queue储存那些可以达到下一个可到达的节点
func canJump2(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	visit := make(map[int]bool)
	queue := make([]int, 0)
	queue = append(queue, len(nums)-1)
	visit[len(nums)-1] = true
	for len(queue) > 0 {
		current := queue[0]
		if current == 0 {
			return true
		}
		queue = queue[1:]

		for i := current - 1; i >= 0; i-- {
			if i+nums[i] >= current && !visit[i] {
				queue = append(queue, i)
				visit[i] = true
			}
		}
	}
	return false
}
