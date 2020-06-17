package search

import (
	"testing"
)

// 在一个 N × N 的方形网格中，每个单元格有两种状态：空（0）或者阻塞（1）
// 返回从左上角到右下角最短路径长度
// https://leetcode-cn.com/problems/shortest-path-in-binary-matrix/ medium
func Test_ShortestPathBinaryMatrix(t *testing.T) {
	//grid := [][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}
	grid := [][]int{{0}}
	t.Log(shortestPathBinaryMatrix(grid))
}

// 最短路径问题，通过用BFS解决
// 广度优先第一次搜索到右下角是为最短路径，所有的路径搜索完仍没到右下角即不存在路径
// 使用队列实现BFS
func shortestPathBinaryMatrix(grid [][]int) int {
	if len(grid) == 0 {
		return -1
	}
	if len(grid) == 1 && len(grid[0]) == 1 && grid[0][0] == 0 {
		return 1
	}
	maxX := len(grid) - 1
	maxY := len(grid[maxX]) - 1
	if grid[0][0] == 1 || grid[len(grid)-1][len(grid[len(grid)-1])-1] == 1 { // 左上角或者右下角为1阻塞
		return -1
	}
	queue := make([]*qValue, 0)
	grid[0][0] = 1 // 置为1
	queue = append(queue, &qValue{
		x:   0,
		y:   0,
		val: grid[0][0],
	})
	// 搜索八个方向
	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]
		x, y := current.x, current.y
		if (x+1 == maxX && y == maxY) || (x == maxX && y+1 == maxY) || (x+1 == maxX && y+1 == maxY) {
			grid[maxX][maxY] = current.val + 1
			return current.val + 1
		}
		if x-1 >= 0 && y-1 >= 0 && grid[x-1][y-1] == 0 {
			grid[x-1][y-1] = current.val + 1
			queue = append(queue, &qValue{
				x:   x - 1,
				y:   y - 1,
				val: grid[x-1][y-1],
			})
		}
		if x-1 >= 0 && grid[x-1][y] == 0 {
			grid[x-1][y] = current.val + 1
			queue = append(queue, &qValue{
				x:   x - 1,
				y:   y,
				val: grid[x-1][y],
			})
		}
		if x-1 >= 0 && y+1 <= maxY && grid[x-1][y+1] == 0 {
			grid[x-1][y+1] = current.val + 1
			queue = append(queue, &qValue{
				x:   x - 1,
				y:   y + 1,
				val: current.val + 1,
			})
		}
		if y-1 >= 0 && grid[x][y-1] == 0 {
			grid[x][y-1] = current.val + 1
			queue = append(queue, &qValue{
				x:   x,
				y:   y - 1,
				val: current.val + 1,
			})
		}
		if y+1 <= maxY && grid[x][y+1] == 0 {
			grid[x][y+1] = current.val + 1
			queue = append(queue, &qValue{
				x:   x,
				y:   y + 1,
				val: current.val + 1,
			})
		}
		if x+1 <= maxX && y-1 >= 0 && grid[x+1][y-1] == 0 {
			grid[x+1][y-1] = current.val + 1
			queue = append(queue, &qValue{
				x:   x + 1,
				y:   y - 1,
				val: current.val + 1,
			})
		}
		if x+1 <= maxX && grid[x+1][y] == 0 {
			grid[x+1][y] = current.val + 1
			queue = append(queue, &qValue{
				x:   x + 1,
				y:   y,
				val: current.val + 1,
			})
		}
		if x+1 <= maxX && y+1 <= maxY && grid[x+1][y+1] == 0 {
			grid[x+1][y+1] = current.val + 1
			queue = append(queue, &qValue{
				x:   x + 1,
				y:   y + 1,
				val: current.val + 1,
			})
		}
	}
	return -1
}

type qValue struct {
	x   int // 横坐标
	y   int // 纵坐标
	val int // 值
}
