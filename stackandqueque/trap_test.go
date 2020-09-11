package stackandqueque

import (
	"github.com/wujiang1994/go-leetcode/common"
	"testing"
)

// 42. 接雨水
// https://leetcode-cn.com/problems/trapping-rain-water/
func Test_Trap(t *testing.T) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	height = []int{5, 4, 1, 2}
	//height = []int{5, 2, 1, 2, 1, 5}
	t.Log(trap(height))
}

// 单调栈实现
// 通过单调递减栈发现右侧有一个比栈顶元素高的时候即可接到雨水
func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	ret := 0
	stack := make([]int, 0)
	for i := 0; i < len(height); i++ {
		// 可省略，最后一步会执行
		//if len(stack) == 0 || height[i] <= height[stack[len(stack)-1]] {
		//	stack = append(stack, i)
		//	continue
		//}
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {

			index := len(stack) - 1
			current := stack[index]
			stack = stack[:index]
			if len(stack) == 0 {
				break
			}
			wide := i - stack[len(stack)-1] - 1
			ret += wide * (common.Min(height[stack[len(stack)-1]], height[i]) - height[current])
		}
		stack = append(stack, i)
	}
	return ret
}
