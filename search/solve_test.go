package search

import (
	"testing"
)

// 给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。
// 找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充
// https://leetcode-cn.com/problems/surrounded-regions/
func Test_Solve(t *testing.T) {
	/*
		board := [][]byte{
			{'X', 'X', 'X', 'X'},
			{'X', 'O', 'O', 'X'},
			{'X', 'X', 'O', 'X'},
			{'X', 'O', 'X', 'X'},
		}
	*/
	board := [][]byte{
		{'O', 'O', 'O'},
		{'O', 'O', 'O'},
		{'O', 'O', 'O'},
	}
	solve(board)
	t.Log(board)
}

// 广度搜索优先，起点为所有外围一圈的O，搜索与其相邻的O，剩余其他的位置都是X
func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	starts := make([][]int, 0)
	visit := make([][]bool, len(board))
	for i := range visit {
		visit[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if i == 0 || i == len(board)-1 {
				if board[i][j] == 'O' {
					starts = append(starts, []int{i, j})
				}
			} else {
				if j == 0 || j == len(board[0])-1 {
					if board[i][j] == 'O' {
						starts = append(starts, []int{i, j})
					}
				}
			}
		}
	}
	for len(starts) != 0 {
		cur := starts[0]
		starts = starts[1:]
		visit[cur[0]][cur[1]] = true
		for _, v := range validNext(board, cur[0], cur[1], visit) {
			visit[v[0]][v[1]] = true
			starts = append(starts, v)
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if visit[i][j] {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}
	return
}

func validNext(board [][]byte, x, y int, visit [][]bool) [][]int {
	ret := make([][]int, 0)
	if x-1 >= 0 && board[x-1][y] == 'O' && !visit[x-1][y] {
		ret = append(ret, []int{x - 1, y})
	}
	if x+1 < len(board) && board[x+1][y] == 'O' && !visit[x+1][y] {
		ret = append(ret, []int{x + 1, y})
	}
	if y-1 >= 0 && board[x][y-1] == 'O' && !visit[x][y-1] {
		ret = append(ret, []int{x, y - 1})
	}
	if y+1 < len(board[0]) && board[x][y+1] == 'O' && !visit[x][y+1] {
		ret = append(ret, []int{x, y + 1})
	}
	return ret
}
