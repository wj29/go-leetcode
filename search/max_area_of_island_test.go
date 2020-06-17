package search

import (
	"github.com/wujiang1994/go-leetcode/common"
	"testing"
)

// 给定一个包含了一些 0 和 1 的非空二维数组 grid 。
// 一个 岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在水平或者竖直方向上相邻。你可以假设 grid 的四个边缘都被 0（代表水）包围着。
// 找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为 0 。)
// https://leetcode-cn.com/problems/max-area-of-island
func Test_MaxAreaOfIsland(t *testing.T) {
	grid := [][]int{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 1, 1}, {0, 0, 0, 1, 1}}
	t.Log(maxAreaOfIsland(grid))
}

// 最大面积问题，通过DFS解决
// 使用栈实现DFS
func maxAreaOfIsland(grid [][]int) int {
	maxX := len(grid) - 1
	maxY := len(grid[maxX]) - 1
	ret := 0
	for i := 0; i <= maxX; i++ {
		for j := 0; j <= maxY; j++ {
			if grid[i][j] == 0 {
				continue
			}
			stack := make([]*sValue, 0)
			grid[i][j] = 0 // 置为0
			stack = append(stack, &sValue{
				x: i,
				y: j,
			})         // 入栈
			count := 0 // 此栈的大小，即面积大小
			for len(stack) != 0 {
				index := len(stack) - 1 // 栈顶
				current := stack[index]
				stack = stack[:index] // 出栈
				x, y := current.x, current.y
				count++
				if x-1 >= 0 && grid[x-1][y] != 0 {
					grid[x-1][y] = 0
					stack = append(stack, &sValue{
						x: x - 1,
						y: y,
					})
				}
				if x+1 <= maxX && grid[x+1][y] != 0 {
					grid[x+1][y] = 0
					stack = append(stack, &sValue{
						x: x + 1,
						y: y,
					})
				}
				if y-1 >= 0 && grid[x][y-1] != 0 {
					grid[x][y-1] = 0
					stack = append(stack, &sValue{
						x: x,
						y: y - 1,
					})
				}
				if y+1 <= maxY && grid[x][y+1] != 0 {
					grid[x][y+1] = 0
					stack = append(stack, &sValue{
						x: x,
						y: y + 1,
					})
				}
			}
			ret = common.Max(ret, count)
		}
	}
	return ret
}

type sValue struct {
	x int
	y int
}
