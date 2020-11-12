package search

import "testing"

// 463. 岛屿的周长
// https://leetcode-cn.com/problems/island-perimeter/
func Test_IslandPerimeter(t *testing.T) {
    grid := [][]int{
        {0, 1, 0, 0},
        {1, 1, 1, 0},
        {0, 1, 0, 0},
        {1, 1, 0, 0},
    }
    t.Log(islandPerimeter(grid))
}

// 对于一个陆地格子的每条边，它被算作岛屿的周长当且仅当这条边为网格的边界或者相邻的另一个格子为水域
// 因此，我们可以遍历每个陆地格子，看其四个方向是否为边界或者水域，如果是，将这条边的贡献加入答案
func islandPerimeter(grid [][]int) int {
    ret := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == 0 {
                continue
            }
            if i == 0 || grid[i-1][j] == 0 {
                ret++
            }
            if i == len(grid)-1 || grid[i+1][j] == 0 {
                ret++
            }
            if j == 0 || grid[i][j-1] == 0 {
                ret++
            }
            if j == len(grid[i])-1 || grid[i][j+1] == 0 {
                ret++
            }
        }
    }
    return ret
}
