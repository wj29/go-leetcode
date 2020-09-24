package backtracking

import (
	"testing"
)

// 36. 有效的数独
// https://leetcode-cn.com/problems/valid-sudoku/
// 37. 解数独
// https://leetcode-cn.com/problems/sudoku-solver/
func Test_SolveSudoku(t *testing.T) {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	t.Log(isValidSudoku(board))
	solveSudoku(board)
	t.Log(board)
}

func solveSudoku(board [][]byte) {
	prev := make([][]byte, len(board))
	for i := range prev {
		prev[i] = make([]byte, len(board[i]))
		copy(prev[i], board[i])
	}
	count := 0
	occupy := make(map[int]map[sudoku][]byte)
	for i := 0; i < len(board); i++ {
		occupy[i] = make(map[sudoku][]byte)
		occupy[i][xSudoku] = make([]byte, 9)
		occupy[i][ySudoku] = make([]byte, 9)
		occupy[i][qSudoku] = make([]byte, 9)
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				count++
				continue
			}
			location := sudokuLocation(i, j)
			for k, v := range location {
				occupy[v][k][int(board[i][j]-'0')-1] = board[i][j]
			}
		}
	}
	board = board[:0]
	backTrackForSolveSudoku(&board, prev, occupy, 0, 0, count)
}

func backTrackForSolveSudoku(board *[][]byte, prev [][]byte, occupy map[int]map[sudoku][]byte, x, y, count int) {
	if count == 0 {
		for i := range prev {
			tmp := make([]byte, len(prev[i]))
			copy(tmp, prev[i])
			*board = append(*board, tmp)
		}
		return
	}
loop:
	for i := '1'; i <= '9'; i++ {
		nextX, nextY := nextIndex(x, y, prev)
		if nextX < 0 && nextY < 0 {
			continue loop
		}
		location := sudokuLocation(nextX, nextY)
		for k, v := range location {
			if occupy[v][k][int(i-'0')-1] == byte(i) {
				continue loop
			}
		}
		prev[nextX][nextY] = byte(i)
		for k, v := range location {
			occupy[v][k][int(i-'0')-1] = byte(i)
		}
		backTrackForSolveSudoku(board, prev, occupy, nextX, nextY, count-1)
		prev[nextX][nextY] = '.'
		for k, v := range location {
			occupy[v][k][int(i-'0')-1] = 0
		}
	}
}

func nextIndex(i, j int, board [][]byte) (int, int) {
	for board[i][j] != '.' {
		j += 1
		if j == 9 {
			i += 1
			j = 0
		}
		if i == 9 && j == 0 {
			return -1, -1
		}
	}
	return i, j
}

func isValidSudoku(board [][]byte) bool {
	occupy := make(map[int]map[sudoku][]byte)
	for i := 0; i < len(board); i++ {
		occupy[i] = make(map[sudoku][]byte)
		occupy[i][xSudoku] = make([]byte, 9)
		occupy[i][ySudoku] = make([]byte, 9)
		occupy[i][qSudoku] = make([]byte, 9)
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				continue
			}
			location := sudokuLocation(i, j)
			for k, v := range location {
				if occupy[v][k][int(board[i][j]-'0')-1] == board[i][j] {
					return false
				}
				occupy[v][k][int(board[i][j]-'0')-1] = board[i][j]
			}
		}
	}
	return true
}

type sudoku int

const ( // 横竖圈占位
	xSudoku sudoku = iota
	ySudoku
	qSudoku
)

func sudokuLocation(x, y int) map[sudoku]int {
	ret := make(map[sudoku]int)
	ret[xSudoku] = x
	ret[ySudoku] = y
	switch {
	case x < 3 && y < 3:
		ret[qSudoku] = 0
	case x < 3 && y >= 3 && y < 6:
		ret[qSudoku] = 1
	case x < 3 && y >= 6 && y < 9:
		ret[qSudoku] = 2
	case x >= 3 && x < 6 && y < 3:
		ret[qSudoku] = 3
	case x >= 3 && x < 6 && y >= 3 && y < 6:
		ret[qSudoku] = 4
	case x >= 3 && x < 6 && y >= 6 && y < 9:
		ret[qSudoku] = 5
	case x >= 6 && x < 9 && y < 3:
		ret[qSudoku] = 6
	case x >= 6 && x < 9 && y >= 3 && y < 6:
		ret[qSudoku] = 7
	case x >= 6 && x < 9 && y >= 6 && y < 9:
		ret[qSudoku] = 8
	}
	return ret
}
