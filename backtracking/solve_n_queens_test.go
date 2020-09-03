package backtracking

import (
	"testing"
)

// 51. N 皇后
// https://leetcode-cn.com/problems/n-queens/
// 52. N皇后 II
// https://leetcode-cn.com/problems/n-queens-ii/
func Test_SolveNQueens(t *testing.T) {
	n := 4
	t.Log(solveNQueens(n))
}

type bit int // 标志横竖斜被占用

const (
	xBit  bit = iota // 0-n-1
	yBit             // 0-n-1
	xzBit            // '/' 左上角为0，向右下依次递增 // 0-2(n-1)
	yzBit            // '\' 右上角为0，向左下依次递增 // 0-2(n-1)
)

func solveNQueens(n int) [][]string {
	prev := make([][]string, n) // 初始化棋盘
	for i := range prev {
		prev[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			prev[i][j] = "."
		}
	}
	ret := make([][]string, 0) // 结果
	for i := range ret {
		ret[i] = make([]string, 0)
	}

	occupy := make(map[int]map[bit]bool) // 占位map，用来判断该location是否能用
	for i := 0; i < 2*n-1; i++ {
		occupy[i] = make(map[bit]bool)
	}
	backTrackForSolveNQueens(&ret, prev, occupy, 0)
	return ret
}

func backTrackForSolveNQueens(ret *[][]string, prev [][]string, occupy map[int]map[bit]bool, start int) {
	if start == len(prev) { // 棋盘已经走完，得到结果
		tmp := make([]string, 0)
		for i := range prev {
			str := ""
			for j := range prev[i] {
				str += prev[i][j]
			}
			tmp = append(tmp, str)
		}
		*ret = append(*ret, tmp)
		return
	}
loop:
	for i := 0; i < len(prev); i++ {
		locationMap := toLocation(start, i, len(prev)) // 计算所占横竖斜位置
		for k, v := range locationMap {
			if occupy[v][k] {
				continue loop
			}
		}
		prev[start][i] = "Q" // 可用置为Q并使该位置横竖斜标记为已占用
		for k, v := range locationMap {
			occupy[v][k] = true
		}
		backTrackForSolveNQueens(ret, prev, occupy, start+1)
		prev[start][i] = "." // 重置状态
		for k, v := range locationMap {
			occupy[v][k] = false
		}
	}
}

func toLocation(x, y, n int) map[bit]int {
	ret := make(map[bit]int)
	ret[xBit] = x
	ret[yBit] = y
	ret[xzBit] = x + y
	ret[yzBit] = n - 1 - y + x
	return ret
}
