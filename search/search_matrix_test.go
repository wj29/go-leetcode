package search

import "testing"

// 74. 搜索二维矩阵
// https://leetcode-cn.com/problems/search-a-2d-matrix/
func Test_SearchMatrix(t *testing.T) {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	target := 3
	t.Log(searchMatrix(matrix, target))
}

// 选择右上角和左下角作为边界值，可以明确搜索方向，时间复杂度O(m+n)
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	index := 0
	for index < len(matrix) {
		if target == matrix[index][len(matrix[index])-1] {
			return true
		} else if target > matrix[index][len(matrix[index])-1] {
			index++
		} else {
			break
		}
	}
	if index == len(matrix) {
		return false
	}
	for _, v := range matrix[index] {
		if target == v {
			return true
		}
	}
	return false
}
