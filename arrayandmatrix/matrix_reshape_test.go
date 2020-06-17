package arrayandmatrix

import "testing"

// 在MATLAB中，有一个非常有用的函数 reshape，它可以将一个矩阵重塑为另一个大小不同的新矩阵，但保留其原始数据
// 给出一个由二维数组表示的矩阵，以及两个正整数r和c，分别表示想要的重构的矩阵的行数和列数
// 重构后的矩阵需要将原始矩阵的所有元素以相同的行遍历顺序填充
// 如果具有给定参数的reshape操作是可行且合理的，则输出新的重塑矩阵；否则，输出原始矩阵
// https://leetcode-cn.com/problems/reshape-the-matrix easy
func Test_MatrixReshape(t *testing.T) {
	nums := [][]int{{1, 2}, {3, 4}}
	r, c := 1, 4
	t.Log(matrixReshape(nums, r, c))
}

func matrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums)*len(nums[0]) != r*c {
		return nums
	}
	ret := make([][]int, r)
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			ret[count] = append(ret[count], nums[i][j])
			if len(ret[count]) >= c {
				count++
			}
		}
	}
	return ret
}
