package dynamicprogramming

import (
	"github.com/wujiang1994/go-leetcode/common"
	"testing"
)

// 在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。
// https://leetcode-cn.com/problems/maximal-square/
//
// 统计并返回其中完全由 1 组成的 正方形 子矩阵的个数。
// https://leetcode-cn.com/problems/count-square-submatrices-with-all-ones/
func Test_MaximalSquare(t *testing.T) {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	t.Log(maximalSquare(matrix))
}

// https://leetcode-cn.com/problems/maximal-square/comments/
// 当我们判断以某个点为正方形右下角时最大的正方形时，那它的上方，左方和左上方三个点也一定是某个正方形的右下角，否则该点为右下角的正方形最大就是它自己了。
// 这是定性的判断，那具体的最大正方形边长呢？我们知道，该点为右下角的正方形的最大边长，最多比它的上方，左方和左上方为右下角的正方形的边长多1，
// 最好的情况是是它的上方，左方和左上方为右下角的正方形的大小都一样的，这样加上该点就可以构成一个更大的正方形。
// 但如果它的上方，左方和左上方为右下角的正方形的大小不一样，合起来就会缺了某个角落，这时候只能取那三个正方形中最小的正方形的边长加1了。

// dp[i][j]表示以matrix[i][j]为右下角的最大正方形的边长
// 那么在matrix[i][j] = 1时 dp[i][j] = {
// 										matrix[i][j] (i=0 || j=0)
// 										min(dp[i-1][j],dp[i][j-1],dp[i-1][j-1]) + 1
//									  }
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	dp := make([][]int, len(matrix))
	ret := 0
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if i == 0 || j == 0 {
				dp[i][j] = int(matrix[i][j] - '0')
			} else if matrix[i][j] == '1' {
				dp[i][j] = common.Min(common.Min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
			if dp[i][j] > ret {
				ret = dp[i][j]
			}
		}
	}
	return ret * ret
}

// 因为dp[i][j]是包含了右下角的最大正方形边长，当dp[i][j]不是0时说明了有以matrix[i][j]为右下角的正方形，其个数等于dp[i][j]
func countSquares(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	dp := make([][]int, len(matrix))
	ret := 0
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if i == 0 || j == 0 {
				dp[i][j] = matrix[i][j]
			} else if matrix[i][j] == 1 {
				dp[i][j] = common.Min(common.Min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
			ret += dp[i][j]
		}
	}
	return ret
}
