package doublepoint

import (
	"testing"
	"github.com/wujiang1994/go-leetcode/common"
)

// 给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai)
// 在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水
// https://leetcode-cn.com/problems/container-with-most-water  medium
func Test_MaxArea(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	t.Log(maxArea(height))
}

// 两层for循环找到每一个组成的面积组合，找到最大的值

// 双指针法，参考官方题解及证明
// https://leetcode-cn.com/problems/container-with-most-water/solution/sheng-zui-duo-shui-de-rong-qi-by-leetcode-solution/
func maxArea(height []int) int {
	ret := 0
	l := 0
	r := len(height) - 1

	for l <= r {
		ret = common.Max(ret, common.Min(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
		} else if height[l] > height[r] {
			r--
		} else {
			l++
			r--
		}
	}
	return ret
}

