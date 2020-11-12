package search

import (
    "testing"
)

// 200. 岛屿数量
// https://leetcode-cn.com/problems/number-of-islands/
func Test_NumIslands(t *testing.T) {
    grid := [][]byte{
        {'1', '1', '0', '0', '0'},
        {'1', '1', '0', '0', '0'},
        {'0', '0', '1', '0', '0'},
        {'0', '0', '0', '1', '1'},
    }
    t.Log(NumIslands(grid))
}

// dfs
func NumIslands(grid [][]byte) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }
    ret := 0
    stack := make([][]int, 0)
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == '0' || grid[i][j] == '2' {
                continue
            }
            stack = append(stack, []int{i, j})
            for len(stack) > 0 {
                current := stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                x, y := current[0], current[1]
                grid[x][y] = '2'
                if x > 0 && grid[x-1][y] == '1' {
                    stack = append(stack, []int{x - 1, y})
                }
                if x < len(grid)-1 && grid[x+1][y] == '1' {
                    stack = append(stack, []int{x + 1, y})
                }
                if y > 0 && grid[x][y-1] == '1' {
                    stack = append(stack, []int{x, y - 1})
                }
                if y < len(grid[x])-1 && grid[x][y+1] == '1' {
                    stack = append(stack, []int{x, y + 1})
                }
            }
            ret++
        }
    }
    return ret
}