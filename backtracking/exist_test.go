package backtracking

import (
	"testing"
)

// 给定一个二维网格和一个单词，找出该单词是否存在于网格中。
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用

// https://leetcode-cn.com/problems/word-search

func Test_Exist(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'V', 'D', 'E', 'E'},
	}
	word := "ABCB"
	t.Log(exist(board, word))
}

// 回溯，找到第一个字符所在位置，上下左右搜索即可
func exist(board [][]byte, word string) bool {
	if len(word) == 0 || len(board) == 0 {
		return false
	}
	starts := make([]*bV, 0)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				starts = append(starts, &bV{x: i, y: j, v: board[i][j]})
			}
		}
	}
	flag := false
	visit := make([][]bool, len(board))
	for i := range visit {
		visit[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < len(starts); i++ {
		visit[starts[i].x][starts[i].y] = true
		backTrackForExist(board, starts[i], word[1:], &flag, visit)
		visit[starts[i].x][starts[i].y] = false
	}
	return flag
}

func backTrackForExist(board [][]byte, start *bV, word string, flag *bool, visit [][]bool) {
	if len(word) == 0 {
		*flag = true
		return
	}
	if *flag {
		return
	}
	validNextIndex := validNext(start.x, start.y, len(board), len(board[0]), board, visit)
	for i := 0; i < len(validNextIndex); i++ {
		if validNextIndex[i].v != word[0] {
			continue
		}
		visit[validNextIndex[i].x][validNextIndex[i].y] = true
		backTrackForExist(board, validNextIndex[i], word[1:], flag, visit)
		visit[validNextIndex[i].x][validNextIndex[i].y] = false
	}
	return
}

func validNext(x, y, maxX, maxY int, board [][]byte, visit [][]bool) []*bV {
	ret := make([]*bV, 0)
	if x-1 >= 0 && !visit[x-1][y] {
		ret = append(ret, &bV{x - 1, y, board[x-1][y]})
	}
	if x+1 < maxX && !visit[x+1][y] {
		ret = append(ret, &bV{x + 1, y, board[x+1][y]})
	}
	if y-1 >= 0 && !visit[x][y-1] {
		ret = append(ret, &bV{x, y - 1, board[x][y-1]})
	}
	if y+1 < maxY && !visit[x][y+1] {
		ret = append(ret, &bV{x, y + 1, board[x][y+1]})
	}
	return ret
}

type bV struct {
	x int
	y int
	v byte
}
