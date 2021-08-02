package arrayandmatrix

import "testing"

// 73. 矩阵置零
// https://leetcode-cn.com/problems/set-matrix-zeroes/
func Test_SetZeroes(t *testing.T) {
    matrix := [][]int{
        {0, 1},
    }
    setZeroes(matrix)
    t.Log(matrix)
}

// 使用首行首列储存0的信息
// 遍历非首行首列，如果matrix[i][j]为0，那么该行该列都将为0，将matrix[i][0]和matrix[0][j]置为0记录下信息
// 此时无法知道首行首列的0的信息，可以在开始时利用两个值记录下0的信息
func setZeroes(matrix [][]int) {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return
    }
    row, col := false, false
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[0]); j++ {
            if matrix[i][j] == 0 {
                matrix[i][0] = 0
                matrix[0][j] = 0
                if i == 0 {
                    row = true
                }
                if j == 0 {
                    col = true
                }
            }
        }
    }

    for i := 1; i < len(matrix); i++ {
        for j := 1; j < len(matrix[0]); j++ {
            if matrix[i][0] == 0 || matrix[0][j] == 0 {
                matrix[i][j] = 0
            }
        }
    }

    if row {
        for i := 0; i < len(matrix[0]); i++ {
            matrix[0][i] = 0
        }
    }
    if col {
        for i := 0; i < len(matrix); i++ {
            matrix[i][0] = 0
        }
    }
    return
}
