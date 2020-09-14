package arrayandmatrix

import (
	"testing"
)

// 给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素
// 54. 螺旋矩阵
// https://leetcode-cn.com/problems/spiral-matrix/
// 59. 螺旋矩阵 II
// https://leetcode-cn.com/problems/spiral-matrix-ii/
func Test_SpiralOrder(t *testing.T) {
	//matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	//t.Log(spiralOrder(matrix))
	n := 3
	t.Log(generateMatrix(n))
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	ret := make([]int, 0)
	spiral(matrix, &ret)
	return ret
}

func spiral(matrix [][]int, ret *[]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	for i := 0; i < len(matrix[0]); i++ {
		*ret = append(*ret, matrix[0][i])
	}
	matrix = matrix[1:]
	for i := 0; i < len(matrix); i++ {
		if len(matrix[i]) > 0 {
			*ret = append(*ret, matrix[i][len(matrix[i])-1])
			matrix[i] = matrix[i][:len(matrix[i])-1]
		}
	}
	if len(matrix) > 0 {
		tmp := matrix[len(matrix)-1]
		matrix = matrix[:len(matrix)-1]
		for i := len(tmp) - 1; i >= 0; i-- {
			*ret = append(*ret, tmp[i])
		}
	}
	for i := len(matrix) - 1; i >= 0; i-- {
		if len(matrix[i]) > 0 {
			*ret = append(*ret, matrix[i][0])
			matrix[i] = matrix[i][1:]
		}
	}
	spiral(matrix, ret)
}

// 设定边界
func generateMatrix(n int) [][]int {
	num := 1
	ret := make([][]int, n)
	for i := range ret {
		ret[i] = make([]int, n)
	}
	up, down, left, right := 0, n-1, 0, n-1
	for num <= n*n {
		for i := left; i <= right; i++ {
			ret[up][i] = num
			num++
		}
		up++
		for i := up; i <= down; i++ {
			ret[i][right] = num
			num++
		}
		right--
		for i := right; i >= left; i-- {
			ret[down][i] = num
			num++
		}
		down--
		for i := down; i >= up; i-- {
			ret[i][left] = num
			num++
		}
		left++
	}
	return ret
}
