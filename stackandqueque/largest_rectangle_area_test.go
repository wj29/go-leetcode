package stackandqueque

import (
	"github.com/wujiang1994/go-leetcode/common"
	"testing"
)

// 84. 柱状图中最大的矩形
// https://leetcode-cn.com/problems/largest-rectangle-in-histogram/
func Test_LargestRectangleArea(t *testing.T) {
	heights := []int{2, 1, 5, 6, 2, 3}
	t.Log(largestRectangleArea(heights))
}

// 单调递增栈
// 对于每一块柱形，以它的高度所能形成的矩形面积
// 其左侧边沿为左侧第一块比它矮的柱形，右侧边沿为右侧第一块比它矮的柱形，其差为宽，可以得到面积
// 对于所有柱形的面积选取最大值即是最大的矩形
// https://blog.csdn.net/Zolewit/article/details/88863970
func largestRectangleArea(heights []int) int {
	stack := make([]int, 0)
	stack = append(stack, -1) // 构建最左侧下标-1计算宽度
	heights = append(heights, 0) // 最右侧一个高度为0的柱形保证全部递增时可以正确结算
	ret := 0
	for i := 0; i < len(heights); i++ {
		for len(stack) > 1 && heights[i] < heights[stack[len(stack)-1]] { // 第i个柱形高度小于前面的高度，即右侧边沿
			index := len(stack) - 1
			current := stack[index]
			stack = stack[:index] // 出栈当前需要计算的柱形
			ret = common.Max(ret, heights[current]*(i-stack[len(stack)-1]-1)) // 左侧边沿
		}
		stack = append(stack, i)
	}
	return ret
}
