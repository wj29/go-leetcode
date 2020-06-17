package arrayandmatrix

import "testing"

// 编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：
// 每行的元素从左到右升序排列
// 每列的元素从上到下升序排列
// https://leetcode-cn.com/problems/search-a-2d-matrix-ii easy
func Test_SearchMatrix(t *testing.T) {
	matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	matrix = [][]int{}
	target := 1
	t.Log(searchMatrix(matrix, target))
	// 输入空数组不会走到循环中 leetcode中测试用例走到循环中panic了 不明所以
}

// target不会出现在行首比他大的行中
// target不会出现在行中比他大的数组后面
func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			if matrix[i][j] > target {
				break
			}
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}

// 时间复杂度可以达到m+n
// TODO